// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 0.2.0
// Version Date: 2020-10-24

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	pb "github.com/techxmind/logserver/interface-defs"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	SubmitSingleEndpoint   endpoint.Endpoint
	SubmitMultipleEndpoint endpoint.Endpoint
	PingEndpoint           endpoint.Endpoint
}

// Endpoints

func (e Endpoints) SubmitSingle(ctx context.Context, in *pb.EventLog) (*pb.Response, error) {
	response, err := e.SubmitSingleEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}

func (e Endpoints) SubmitMultiple(ctx context.Context, in *pb.EventLogs) (*pb.Response, error) {
	response, err := e.SubmitMultipleEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}

func (e Endpoints) Ping(ctx context.Context, in *pb.Empty) (*pb.Response, error) {
	response, err := e.PingEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}

// Make Endpoints

func MakeSubmitSingleEndpoint(s pb.LogServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.EventLog)
		v, err := s.SubmitSingle(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeSubmitMultipleEndpoint(s pb.LogServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.EventLogs)
		v, err := s.SubmitMultiple(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakePingEndpoint(s pb.LogServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.Empty)
		v, err := s.Ping(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"SubmitSingle":   {},
		"SubmitMultiple": {},
		"Ping":           {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "SubmitSingle" {
			e.SubmitSingleEndpoint = middleware(e.SubmitSingleEndpoint)
		}
		if inc == "SubmitMultiple" {
			e.SubmitMultipleEndpoint = middleware(e.SubmitMultipleEndpoint)
		}
		if inc == "Ping" {
			e.PingEndpoint = middleware(e.PingEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"SubmitSingle":   {},
		"SubmitMultiple": {},
		"Ping":           {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "SubmitSingle" {
			e.SubmitSingleEndpoint = middleware("SubmitSingle", e.SubmitSingleEndpoint)
		}
		if inc == "SubmitMultiple" {
			e.SubmitMultipleEndpoint = middleware("SubmitMultiple", e.SubmitMultipleEndpoint)
		}
		if inc == "Ping" {
			e.PingEndpoint = middleware("Ping", e.PingEndpoint)
		}
	}
}
