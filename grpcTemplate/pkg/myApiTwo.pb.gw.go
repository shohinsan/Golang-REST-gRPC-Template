// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: myApiTwo.proto

/*
Package pkg is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package pkg

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_MySecondApi_ThirdGetRpc_0(ctx context.Context, marshaler runtime.Marshaler, client MySecondApiClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := client.ThirdGetRpc(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_MySecondApi_ThirdGetRpc_0(ctx context.Context, marshaler runtime.Marshaler, server MySecondApiServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := server.ThirdGetRpc(ctx, &protoReq)
	return msg, metadata, err

}

func request_MySecondApi_FourthGetRpc_0(ctx context.Context, marshaler runtime.Marshaler, client MySecondApiClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := client.FourthGetRpc(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_MySecondApi_FourthGetRpc_0(ctx context.Context, marshaler runtime.Marshaler, server MySecondApiServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := server.FourthGetRpc(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterMySecondApiHandlerServer registers the http handlers for service MySecondApi to "mux".
// UnaryRPC     :call MySecondApiServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterMySecondApiHandlerFromEndpoint instead.
func RegisterMySecondApiHandlerServer(ctx context.Context, mux *runtime.ServeMux, server MySecondApiServer) error {

	mux.Handle("GET", pattern_MySecondApi_ThirdGetRpc_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/routes.MySecondApi/ThirdGetRpc", runtime.WithHTTPPathPattern("/third"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_MySecondApi_ThirdGetRpc_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MySecondApi_ThirdGetRpc_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_MySecondApi_FourthGetRpc_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/routes.MySecondApi/FourthGetRpc", runtime.WithHTTPPathPattern("/fourth"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_MySecondApi_FourthGetRpc_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MySecondApi_FourthGetRpc_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterMySecondApiHandlerFromEndpoint is same as RegisterMySecondApiHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterMySecondApiHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterMySecondApiHandler(ctx, mux, conn)
}

// RegisterMySecondApiHandler registers the http handlers for service MySecondApi to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterMySecondApiHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterMySecondApiHandlerClient(ctx, mux, NewMySecondApiClient(conn))
}

// RegisterMySecondApiHandlerClient registers the http handlers for service MySecondApi
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "MySecondApiClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "MySecondApiClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "MySecondApiClient" to call the correct interceptors.
func RegisterMySecondApiHandlerClient(ctx context.Context, mux *runtime.ServeMux, client MySecondApiClient) error {

	mux.Handle("GET", pattern_MySecondApi_ThirdGetRpc_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/routes.MySecondApi/ThirdGetRpc", runtime.WithHTTPPathPattern("/third"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_MySecondApi_ThirdGetRpc_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MySecondApi_ThirdGetRpc_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_MySecondApi_FourthGetRpc_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/routes.MySecondApi/FourthGetRpc", runtime.WithHTTPPathPattern("/fourth"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_MySecondApi_FourthGetRpc_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MySecondApi_FourthGetRpc_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_MySecondApi_ThirdGetRpc_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"third"}, ""))

	pattern_MySecondApi_FourthGetRpc_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"fourth"}, ""))
)

var (
	forward_MySecondApi_ThirdGetRpc_0 = runtime.ForwardResponseMessage

	forward_MySecondApi_FourthGetRpc_0 = runtime.ForwardResponseMessage
)
