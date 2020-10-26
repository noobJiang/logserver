// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 0.2.0
// Version Date: 2020-10-24

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/techxmind/logserver/interface-defs"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC LogServiceServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.LogServiceServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// logservice

		submitsingle: grpctransport.NewServer(
			endpoints.SubmitSingleEndpoint,
			DecodeGRPCSubmitSingleRequest,
			EncodeGRPCSubmitSingleResponse,
			serverOptions...,
		),
		submitmultiple: grpctransport.NewServer(
			endpoints.SubmitMultipleEndpoint,
			DecodeGRPCSubmitMultipleRequest,
			EncodeGRPCSubmitMultipleResponse,
			serverOptions...,
		),
		ping: grpctransport.NewServer(
			endpoints.PingEndpoint,
			DecodeGRPCPingRequest,
			EncodeGRPCPingResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the LogServiceServer interface
type grpcServer struct {
	submitsingle   grpctransport.Handler
	submitmultiple grpctransport.Handler
	ping           grpctransport.Handler
}

// Methods for grpcServer to implement LogServiceServer interface

func (s *grpcServer) SubmitSingle(ctx context.Context, req *pb.EventLog) (*pb.Response, error) {
	_, rep, err := s.submitsingle.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Response), nil
}

func (s *grpcServer) SubmitMultiple(ctx context.Context, req *pb.EventLogs) (*pb.Response, error) {
	_, rep, err := s.submitmultiple.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Response), nil
}

func (s *grpcServer) Ping(ctx context.Context, req *pb.Empty) (*pb.Response, error) {
	_, rep, err := s.ping.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Response), nil
}

// Server Decode

// DecodeGRPCSubmitSingleRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC submitsingle request to a user-domain submitsingle request. Primarily useful in a server.
func DecodeGRPCSubmitSingleRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.EventLog)
	return req, nil
}

// DecodeGRPCSubmitMultipleRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC submitmultiple request to a user-domain submitmultiple request. Primarily useful in a server.
func DecodeGRPCSubmitMultipleRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.EventLogs)
	return req, nil
}

// DecodeGRPCPingRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC ping request to a user-domain ping request. Primarily useful in a server.
func DecodeGRPCPingRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Empty)
	return req, nil
}

// Server Encode

// EncodeGRPCSubmitSingleResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain submitsingle response to a gRPC submitsingle reply. Primarily useful in a server.
func EncodeGRPCSubmitSingleResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.Response)
	return resp, nil
}

// EncodeGRPCSubmitMultipleResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain submitmultiple response to a gRPC submitmultiple reply. Primarily useful in a server.
func EncodeGRPCSubmitMultipleResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.Response)
	return resp, nil
}

// EncodeGRPCPingResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain ping response to a gRPC ping reply. Primarily useful in a server.
func EncodeGRPCPingResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.Response)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
