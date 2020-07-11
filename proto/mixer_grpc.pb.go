// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MixerClient is the client API for Mixer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MixerClient interface {
	// Query DataCommons Graph with Sparql.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// Query DataCommons Graph with Sparql.
	QueryPost(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// Fetch property labels adjacent of nodes
	GetPropertyLabels(ctx context.Context, in *GetPropertyLabelsRequest, opts ...grpc.CallOption) (*GetPropertyLabelsResponse, error)
	// Fetch property labels adjacent of nodes
	GetPropertyLabelsPost(ctx context.Context, in *GetPropertyLabelsRequest, opts ...grpc.CallOption) (*GetPropertyLabelsResponse, error)
	// Fetch nodes that linked to source nodes with a given property.
	GetPropertyValues(ctx context.Context, in *GetPropertyValuesRequest, opts ...grpc.CallOption) (*GetPropertyValuesResponse, error)
	// Fetch nodes that linked to source nodes with a given property.
	GetPropertyValuesPost(ctx context.Context, in *GetPropertyValuesRequest, opts ...grpc.CallOption) (*GetPropertyValuesResponse, error)
	// Fetch triples that have the given nodes as subject or object.
	GetTriples(ctx context.Context, in *GetTriplesRequest, opts ...grpc.CallOption) (*GetTriplesResponse, error)
	// Fetch triples that have the given nodes as subject or object.
	GetTriplesPost(ctx context.Context, in *GetTriplesRequest, opts ...grpc.CallOption) (*GetTriplesResponse, error)
	// Get populations for a list of places, given the population type and constraining property
	// values.
	GetPopulations(ctx context.Context, in *GetPopulationsRequest, opts ...grpc.CallOption) (*GetPopulationsResponse, error)
	// Get observations for a list of population, given the observation constraints.
	GetObservations(ctx context.Context, in *GetObservationsRequest, opts ...grpc.CallOption) (*GetObservationsResponse, error)
	// Get places contained in parent places.
	GetPlacesIn(ctx context.Context, in *GetPlacesInRequest, opts ...grpc.CallOption) (*GetPlacesInResponse, error)
	// Get places contained in parent places.
	GetPlacesInPost(ctx context.Context, in *GetPlacesInRequest, opts ...grpc.CallOption) (*GetPlacesInResponse, error)
	// Get population and observation data for a place.
	GetPopObs(ctx context.Context, in *GetPopObsRequest, opts ...grpc.CallOption) (*GetPopObsResponse, error)
	// Get observation data for a list of places, given place type, population type, and
	// population constraining properties.
	GetPlaceObs(ctx context.Context, in *GetPlaceObsRequest, opts ...grpc.CallOption) (*GetPlaceObsResponse, error)
	// Get stats of places by StatisticalVariable.
	GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error)
	// Get a list of possible population type, measured property, and PVs for a given place type.
	GetPopCategory(ctx context.Context, in *GetPopCategoryRequest, opts ...grpc.CallOption) (*GetPopCategoryResponse, error)
	// Get related places for a given statistical variable.
	GetRelatedPlaces(ctx context.Context, in *GetRelatedPlacesRequest, opts ...grpc.CallOption) (*GetRelatedPlacesResponse, error)
	// Get interesting aspects for places.
	GetInterestingPlaceAspects(ctx context.Context, in *GetInterestingPlaceAspectsRequest, opts ...grpc.CallOption) (*GetInterestingPlaceAspectsResponse, error)
	// Get chart data.
	GetChartData(ctx context.Context, in *GetChartDataRequest, opts ...grpc.CallOption) (*GetChartDataResponse, error)
	// Translate Sparql Query into translation results.
	Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error)
	// Given a text search query, return all entities matching the query.
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	// Give a list of place dcids, return all the statistical variables for each
	// place.
	GetPlaceStatsVar(ctx context.Context, in *GetPlaceStatsVarRequest, opts ...grpc.CallOption) (*GetPlaceStatsVarResponse, error)
}

type mixerClient struct {
	cc grpc.ClientConnInterface
}

func NewMixerClient(cc grpc.ClientConnInterface) MixerClient {
	return &mixerClient{cc}
}

func (c *mixerClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) QueryPost(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/QueryPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetPropertyLabels(ctx context.Context, in *GetPropertyLabelsRequest, opts ...grpc.CallOption) (*GetPropertyLabelsResponse, error) {
	out := new(GetPropertyLabelsResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetPropertyLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetPropertyLabelsPost(ctx context.Context, in *GetPropertyLabelsRequest, opts ...grpc.CallOption) (*GetPropertyLabelsResponse, error) {
	out := new(GetPropertyLabelsResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetPropertyLabelsPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetPropertyValues(ctx context.Context, in *GetPropertyValuesRequest, opts ...grpc.CallOption) (*GetPropertyValuesResponse, error) {
	out := new(GetPropertyValuesResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetPropertyValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetPropertyValuesPost(ctx context.Context, in *GetPropertyValuesRequest, opts ...grpc.CallOption) (*GetPropertyValuesResponse, error) {
	out := new(GetPropertyValuesResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetPropertyValuesPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetTriples(ctx context.Context, in *GetTriplesRequest, opts ...grpc.CallOption) (*GetTriplesResponse, error) {
	out := new(GetTriplesResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetTriples", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetTriplesPost(ctx context.Context, in *GetTriplesRequest, opts ...grpc.CallOption) (*GetTriplesResponse, error) {
	out := new(GetTriplesResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetTriplesPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetPopulations(ctx context.Context, in *GetPopulationsRequest, opts ...grpc.CallOption) (*GetPopulationsResponse, error) {
	out := new(GetPopulationsResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetPopulations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetObservations(ctx context.Context, in *GetObservationsRequest, opts ...grpc.CallOption) (*GetObservationsResponse, error) {
	out := new(GetObservationsResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetObservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetPlacesIn(ctx context.Context, in *GetPlacesInRequest, opts ...grpc.CallOption) (*GetPlacesInResponse, error) {
	out := new(GetPlacesInResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetPlacesIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetPlacesInPost(ctx context.Context, in *GetPlacesInRequest, opts ...grpc.CallOption) (*GetPlacesInResponse, error) {
	out := new(GetPlacesInResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetPlacesInPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetPopObs(ctx context.Context, in *GetPopObsRequest, opts ...grpc.CallOption) (*GetPopObsResponse, error) {
	out := new(GetPopObsResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetPopObs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetPlaceObs(ctx context.Context, in *GetPlaceObsRequest, opts ...grpc.CallOption) (*GetPlaceObsResponse, error) {
	out := new(GetPlaceObsResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetPlaceObs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetPopCategory(ctx context.Context, in *GetPopCategoryRequest, opts ...grpc.CallOption) (*GetPopCategoryResponse, error) {
	out := new(GetPopCategoryResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetPopCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetRelatedPlaces(ctx context.Context, in *GetRelatedPlacesRequest, opts ...grpc.CallOption) (*GetRelatedPlacesResponse, error) {
	out := new(GetRelatedPlacesResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetRelatedPlaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetInterestingPlaceAspects(ctx context.Context, in *GetInterestingPlaceAspectsRequest, opts ...grpc.CallOption) (*GetInterestingPlaceAspectsResponse, error) {
	out := new(GetInterestingPlaceAspectsResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetInterestingPlaceAspects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetChartData(ctx context.Context, in *GetChartDataRequest, opts ...grpc.CallOption) (*GetChartDataResponse, error) {
	out := new(GetChartDataResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetChartData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error) {
	out := new(TranslateResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/Translate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) GetPlaceStatsVar(ctx context.Context, in *GetPlaceStatsVarRequest, opts ...grpc.CallOption) (*GetPlaceStatsVarResponse, error) {
	out := new(GetPlaceStatsVarResponse)
	err := c.cc.Invoke(ctx, "/datacommons.Mixer/GetPlaceStatsVar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MixerServer is the server API for Mixer service.
// All implementations should embed UnimplementedMixerServer
// for forward compatibility
type MixerServer interface {
	// Query DataCommons Graph with Sparql.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	// Query DataCommons Graph with Sparql.
	QueryPost(context.Context, *QueryRequest) (*QueryResponse, error)
	// Fetch property labels adjacent of nodes
	GetPropertyLabels(context.Context, *GetPropertyLabelsRequest) (*GetPropertyLabelsResponse, error)
	// Fetch property labels adjacent of nodes
	GetPropertyLabelsPost(context.Context, *GetPropertyLabelsRequest) (*GetPropertyLabelsResponse, error)
	// Fetch nodes that linked to source nodes with a given property.
	GetPropertyValues(context.Context, *GetPropertyValuesRequest) (*GetPropertyValuesResponse, error)
	// Fetch nodes that linked to source nodes with a given property.
	GetPropertyValuesPost(context.Context, *GetPropertyValuesRequest) (*GetPropertyValuesResponse, error)
	// Fetch triples that have the given nodes as subject or object.
	GetTriples(context.Context, *GetTriplesRequest) (*GetTriplesResponse, error)
	// Fetch triples that have the given nodes as subject or object.
	GetTriplesPost(context.Context, *GetTriplesRequest) (*GetTriplesResponse, error)
	// Get populations for a list of places, given the population type and constraining property
	// values.
	GetPopulations(context.Context, *GetPopulationsRequest) (*GetPopulationsResponse, error)
	// Get observations for a list of population, given the observation constraints.
	GetObservations(context.Context, *GetObservationsRequest) (*GetObservationsResponse, error)
	// Get places contained in parent places.
	GetPlacesIn(context.Context, *GetPlacesInRequest) (*GetPlacesInResponse, error)
	// Get places contained in parent places.
	GetPlacesInPost(context.Context, *GetPlacesInRequest) (*GetPlacesInResponse, error)
	// Get population and observation data for a place.
	GetPopObs(context.Context, *GetPopObsRequest) (*GetPopObsResponse, error)
	// Get observation data for a list of places, given place type, population type, and
	// population constraining properties.
	GetPlaceObs(context.Context, *GetPlaceObsRequest) (*GetPlaceObsResponse, error)
	// Get stats of places by StatisticalVariable.
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
	// Get a list of possible population type, measured property, and PVs for a given place type.
	GetPopCategory(context.Context, *GetPopCategoryRequest) (*GetPopCategoryResponse, error)
	// Get related places for a given statistical variable.
	GetRelatedPlaces(context.Context, *GetRelatedPlacesRequest) (*GetRelatedPlacesResponse, error)
	// Get interesting aspects for places.
	GetInterestingPlaceAspects(context.Context, *GetInterestingPlaceAspectsRequest) (*GetInterestingPlaceAspectsResponse, error)
	// Get chart data.
	GetChartData(context.Context, *GetChartDataRequest) (*GetChartDataResponse, error)
	// Translate Sparql Query into translation results.
	Translate(context.Context, *TranslateRequest) (*TranslateResponse, error)
	// Given a text search query, return all entities matching the query.
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	// Give a list of place dcids, return all the statistical variables for each
	// place.
	GetPlaceStatsVar(context.Context, *GetPlaceStatsVarRequest) (*GetPlaceStatsVarResponse, error)
}

// UnimplementedMixerServer should be embedded to have forward compatible implementations.
type UnimplementedMixerServer struct {
}

func (*UnimplementedMixerServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (*UnimplementedMixerServer) QueryPost(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPost not implemented")
}
func (*UnimplementedMixerServer) GetPropertyLabels(context.Context, *GetPropertyLabelsRequest) (*GetPropertyLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPropertyLabels not implemented")
}
func (*UnimplementedMixerServer) GetPropertyLabelsPost(context.Context, *GetPropertyLabelsRequest) (*GetPropertyLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPropertyLabelsPost not implemented")
}
func (*UnimplementedMixerServer) GetPropertyValues(context.Context, *GetPropertyValuesRequest) (*GetPropertyValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPropertyValues not implemented")
}
func (*UnimplementedMixerServer) GetPropertyValuesPost(context.Context, *GetPropertyValuesRequest) (*GetPropertyValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPropertyValuesPost not implemented")
}
func (*UnimplementedMixerServer) GetTriples(context.Context, *GetTriplesRequest) (*GetTriplesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTriples not implemented")
}
func (*UnimplementedMixerServer) GetTriplesPost(context.Context, *GetTriplesRequest) (*GetTriplesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTriplesPost not implemented")
}
func (*UnimplementedMixerServer) GetPopulations(context.Context, *GetPopulationsRequest) (*GetPopulationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPopulations not implemented")
}
func (*UnimplementedMixerServer) GetObservations(context.Context, *GetObservationsRequest) (*GetObservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObservations not implemented")
}
func (*UnimplementedMixerServer) GetPlacesIn(context.Context, *GetPlacesInRequest) (*GetPlacesInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlacesIn not implemented")
}
func (*UnimplementedMixerServer) GetPlacesInPost(context.Context, *GetPlacesInRequest) (*GetPlacesInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlacesInPost not implemented")
}
func (*UnimplementedMixerServer) GetPopObs(context.Context, *GetPopObsRequest) (*GetPopObsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPopObs not implemented")
}
func (*UnimplementedMixerServer) GetPlaceObs(context.Context, *GetPlaceObsRequest) (*GetPlaceObsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaceObs not implemented")
}
func (*UnimplementedMixerServer) GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (*UnimplementedMixerServer) GetPopCategory(context.Context, *GetPopCategoryRequest) (*GetPopCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPopCategory not implemented")
}
func (*UnimplementedMixerServer) GetRelatedPlaces(context.Context, *GetRelatedPlacesRequest) (*GetRelatedPlacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelatedPlaces not implemented")
}
func (*UnimplementedMixerServer) GetInterestingPlaceAspects(context.Context, *GetInterestingPlaceAspectsRequest) (*GetInterestingPlaceAspectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInterestingPlaceAspects not implemented")
}
func (*UnimplementedMixerServer) GetChartData(context.Context, *GetChartDataRequest) (*GetChartDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChartData not implemented")
}
func (*UnimplementedMixerServer) Translate(context.Context, *TranslateRequest) (*TranslateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Translate not implemented")
}
func (*UnimplementedMixerServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedMixerServer) GetPlaceStatsVar(context.Context, *GetPlaceStatsVarRequest) (*GetPlaceStatsVarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaceStatsVar not implemented")
}

func RegisterMixerServer(s *grpc.Server, srv MixerServer) {
	s.RegisterService(&_Mixer_serviceDesc, srv)
}

func _Mixer_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_QueryPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).QueryPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/QueryPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).QueryPost(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetPropertyLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPropertyLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetPropertyLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetPropertyLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetPropertyLabels(ctx, req.(*GetPropertyLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetPropertyLabelsPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPropertyLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetPropertyLabelsPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetPropertyLabelsPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetPropertyLabelsPost(ctx, req.(*GetPropertyLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetPropertyValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPropertyValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetPropertyValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetPropertyValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetPropertyValues(ctx, req.(*GetPropertyValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetPropertyValuesPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPropertyValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetPropertyValuesPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetPropertyValuesPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetPropertyValuesPost(ctx, req.(*GetPropertyValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetTriples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTriplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetTriples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetTriples",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetTriples(ctx, req.(*GetTriplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetTriplesPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTriplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetTriplesPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetTriplesPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetTriplesPost(ctx, req.(*GetTriplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetPopulations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPopulationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetPopulations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetPopulations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetPopulations(ctx, req.(*GetPopulationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetObservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObservationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetObservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetObservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetObservations(ctx, req.(*GetObservationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetPlacesIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlacesInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetPlacesIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetPlacesIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetPlacesIn(ctx, req.(*GetPlacesInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetPlacesInPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlacesInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetPlacesInPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetPlacesInPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetPlacesInPost(ctx, req.(*GetPlacesInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetPopObs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPopObsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetPopObs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetPopObs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetPopObs(ctx, req.(*GetPopObsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetPlaceObs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlaceObsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetPlaceObs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetPlaceObs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetPlaceObs(ctx, req.(*GetPlaceObsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetPopCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPopCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetPopCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetPopCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetPopCategory(ctx, req.(*GetPopCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetRelatedPlaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelatedPlacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetRelatedPlaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetRelatedPlaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetRelatedPlaces(ctx, req.(*GetRelatedPlacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetInterestingPlaceAspects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInterestingPlaceAspectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetInterestingPlaceAspects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetInterestingPlaceAspects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetInterestingPlaceAspects(ctx, req.(*GetInterestingPlaceAspectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetChartData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChartDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetChartData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetChartData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetChartData(ctx, req.(*GetChartDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/Translate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).Translate(ctx, req.(*TranslateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_GetPlaceStatsVar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlaceStatsVarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).GetPlaceStatsVar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacommons.Mixer/GetPlaceStatsVar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).GetPlaceStatsVar(ctx, req.(*GetPlaceStatsVarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mixer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "datacommons.Mixer",
	HandlerType: (*MixerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Mixer_Query_Handler,
		},
		{
			MethodName: "QueryPost",
			Handler:    _Mixer_QueryPost_Handler,
		},
		{
			MethodName: "GetPropertyLabels",
			Handler:    _Mixer_GetPropertyLabels_Handler,
		},
		{
			MethodName: "GetPropertyLabelsPost",
			Handler:    _Mixer_GetPropertyLabelsPost_Handler,
		},
		{
			MethodName: "GetPropertyValues",
			Handler:    _Mixer_GetPropertyValues_Handler,
		},
		{
			MethodName: "GetPropertyValuesPost",
			Handler:    _Mixer_GetPropertyValuesPost_Handler,
		},
		{
			MethodName: "GetTriples",
			Handler:    _Mixer_GetTriples_Handler,
		},
		{
			MethodName: "GetTriplesPost",
			Handler:    _Mixer_GetTriplesPost_Handler,
		},
		{
			MethodName: "GetPopulations",
			Handler:    _Mixer_GetPopulations_Handler,
		},
		{
			MethodName: "GetObservations",
			Handler:    _Mixer_GetObservations_Handler,
		},
		{
			MethodName: "GetPlacesIn",
			Handler:    _Mixer_GetPlacesIn_Handler,
		},
		{
			MethodName: "GetPlacesInPost",
			Handler:    _Mixer_GetPlacesInPost_Handler,
		},
		{
			MethodName: "GetPopObs",
			Handler:    _Mixer_GetPopObs_Handler,
		},
		{
			MethodName: "GetPlaceObs",
			Handler:    _Mixer_GetPlaceObs_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _Mixer_GetStats_Handler,
		},
		{
			MethodName: "GetPopCategory",
			Handler:    _Mixer_GetPopCategory_Handler,
		},
		{
			MethodName: "GetRelatedPlaces",
			Handler:    _Mixer_GetRelatedPlaces_Handler,
		},
		{
			MethodName: "GetInterestingPlaceAspects",
			Handler:    _Mixer_GetInterestingPlaceAspects_Handler,
		},
		{
			MethodName: "GetChartData",
			Handler:    _Mixer_GetChartData_Handler,
		},
		{
			MethodName: "Translate",
			Handler:    _Mixer_Translate_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Mixer_Search_Handler,
		},
		{
			MethodName: "GetPlaceStatsVar",
			Handler:    _Mixer_GetPlaceStatsVar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mixer.proto",
}