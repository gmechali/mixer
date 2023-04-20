// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package server is the main server
package server

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/datacommonsorg/mixer/internal/merger"
	"github.com/datacommonsorg/mixer/internal/server/resource"
	v2 "github.com/datacommonsorg/mixer/internal/server/v2"
	v2p "github.com/datacommonsorg/mixer/internal/server/v2/properties"
	v2pv "github.com/datacommonsorg/mixer/internal/server/v2/propertyvalues"
	"github.com/datacommonsorg/mixer/internal/util"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	pbv2 "github.com/datacommonsorg/mixer/internal/proto/v2"
)

func fetchRemote(
	metadata *resource.Metadata,
	httpClient *http.Client,
	apiPath string,
	in proto.Message,
	out proto.Message,
) error {
	url := metadata.RemoteMixerDomain + apiPath
	jsonValue, err := protojson.Marshal(in)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", metadata.RemoteMixerAPIKey)
	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	// Read response body
	var responseBodyBytes []byte
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("remote mixer response not ok: %s", response.Status)
	}
	responseBodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	// Convert response body to string
	return protojson.Unmarshal(responseBodyBytes, out)
}

// V2Resolve implements API for mixer.V2Resolve.
func (s *Server) V2Resolve(
	ctx context.Context, in *pbv2.ResolveRequest,
) (*pbv2.ResolveResponse, error) {
	resp, err := V2ResolveCore(ctx, s.store, s.mapsClient, in)
	if err != nil {
		return nil, err
	}
	if s.metadata.RemoteMixerDomain != "" {
		remoteResp := &pbv2.ResolveResponse{}
		err := fetchRemote(s.metadata, s.httpClient, "/v2/resolve", in, remoteResp)
		if err != nil {
			return nil, err
		}
		return merger.MergeResolve(resp, remoteResp), nil
	}
	return resp, nil
}

// V2Node implements API for mixer.V2Node.
func (s *Server) V2Node(
	ctx context.Context, in *pbv2.NodeRequest,
) (*pbv2.NodeResponse, error) {
	arcs, err := v2.ParseProperty(in.GetProperty())
	if err != nil {
		return nil, err
	}
	if len(arcs) == 1 {
		arc := arcs[0]
		direction := util.DirectionOut
		if !arc.Out {
			direction = util.DirectionIn
		}

		if arc.SingleProp != "" && arc.Wildcard == "+" {
			// Examples:
			//   <-containedInPlace+{typeOf:City}
			return v2pv.LinkedPropertyValues(
				ctx,
				s.store,
				s.cache,
				in.GetNodes(),
				arc.SingleProp,
				direction,
				arc.Filter,
			)
		}

		if arc.SingleProp == "" && len(arc.BracketProps) == 0 {
			// Examples:
			//   ->
			//   <-
			return v2p.API(ctx, s.store, in.GetNodes(), direction)
		}

		var properties []string
		if arc.SingleProp != "" {
			if arc.Wildcard == "" {
				// Examples:
				//   ->name
				//   <-containedInPlace
				properties = []string{arc.SingleProp}
			}
		} else { // arc.SingleProp == ""
			// Examples:
			//   ->[name, address]
			properties = arc.BracketProps
		}
		if len(properties) > 0 {
			return v2pv.API(
				ctx,
				s.store,
				s.metadata,
				in.GetNodes(),
				properties,
				direction,
				int(in.GetLimit()),
				in.NextToken,
			)
		}
	}
	return nil, nil
}

// V2Event implements API for mixer.V2Event.
func (s *Server) V2Event(
	ctx context.Context, in *pbv2.EventRequest,
) (*pbv2.EventResponse, error) {
	resp, err := V2EventCore(ctx, s.store, in)
	if err != nil {
		return nil, err
	}
	if s.metadata.RemoteMixerDomain != "" {
		remoteResp := &pbv2.EventResponse{}
		err := fetchRemote(s.metadata, s.httpClient, "/v2/event", in, remoteResp)
		if err != nil {
			return nil, err
		}
		return merger.MergeEvent(resp, remoteResp), nil
	}
	return resp, nil
}

// V2Observation implements API for mixer.V2Observation.
func (s *Server) V2Observation(
	ctx context.Context, in *pbv2.ObservationRequest,
) (*pbv2.ObservationResponse, error) {
	resp, err := V2ObservationCore(ctx, s.store, in)
	if err != nil {
		return nil, err
	}
	if s.metadata.RemoteMixerDomain != "" {
		remoteResp := &pbv2.ObservationResponse{}
		err := fetchRemote(s.metadata, s.httpClient, "/v2/observation", in, remoteResp)
		if err != nil {
			return nil, err
		}
		return merger.MergeObservation(resp, remoteResp), nil
	}
	return resp, nil
}
