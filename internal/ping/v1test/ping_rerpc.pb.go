// Code generated by protoc-gen-go-rerpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-rerpc v0.0.1
// - protoc             v3.17.3
// source: internal/ping/v1test/ping.proto

package pingpb

import (
	context "context"
	errors "errors"
	rerpc "github.com/rerpc/rerpc"
	callstream "github.com/rerpc/rerpc/callstream"
	handlerstream "github.com/rerpc/rerpc/handlerstream"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the
// rerpc package are compatible. If you get a compiler error that this constant
// isn't defined, this code was generated with a version of rerpc newer than the
// one compiled into your binary. You can fix the problem by either regenerating
// this code with an older version of rerpc or updating the rerpc version
// compiled into your binary.
const _ = rerpc.SupportsCodeGenV0 // requires reRPC v0.0.1 or later

// PingServiceClientReRPC is a client for the internal.ping.v1test.PingService
// service.
type PingServiceClientReRPC interface {
	Ping(ctx context.Context, req *PingRequest, opts ...rerpc.CallOption) (*PingResponse, error)
	Fail(ctx context.Context, req *FailRequest, opts ...rerpc.CallOption) (*FailResponse, error)
	Sum(ctx context.Context, opts ...rerpc.CallOption) *callstream.Client[SumRequest, SumResponse]
	CountUp(ctx context.Context, req *CountUpRequest, opts ...rerpc.CallOption) (*callstream.Server[CountUpResponse], error)
	CumSum(ctx context.Context, opts ...rerpc.CallOption) *callstream.Bidirectional[CumSumRequest, CumSumResponse]
}

type pingServiceClientReRPC struct {
	doer    rerpc.Doer
	baseURL string
	options []rerpc.CallOption
}

// NewPingServiceClientReRPC constructs a client for the
// internal.ping.v1test.PingService service. Call options passed here apply to
// all calls made with this client.
//
// The URL supplied here should be the base URL for the gRPC server (e.g.,
// https://api.acme.com or https://acme.com/grpc).
func NewPingServiceClientReRPC(baseURL string, doer rerpc.Doer, opts ...rerpc.CallOption) PingServiceClientReRPC {
	return &pingServiceClientReRPC{
		baseURL: strings.TrimRight(baseURL, "/"),
		doer:    doer,
		options: opts,
	}
}

func (c *pingServiceClientReRPC) mergeOptions(opts []rerpc.CallOption) []rerpc.CallOption {
	merged := make([]rerpc.CallOption, 0, len(c.options)+len(opts))
	for _, o := range c.options {
		merged = append(merged, o)
	}
	for _, o := range opts {
		merged = append(merged, o)
	}
	return merged
}

// Ping calls internal.ping.v1test.PingService.Ping. Call options passed here
// apply only to this call.
func (c *pingServiceClientReRPC) Ping(ctx context.Context, req *PingRequest, opts ...rerpc.CallOption) (*PingResponse, error) {
	merged := c.mergeOptions(opts)
	call := rerpc.NewClientFunc[PingRequest, PingResponse](
		c.doer,
		c.baseURL,
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"Ping",                 // protobuf method
		merged...,
	)
	return call(ctx, req)
}

// Fail calls internal.ping.v1test.PingService.Fail. Call options passed here
// apply only to this call.
func (c *pingServiceClientReRPC) Fail(ctx context.Context, req *FailRequest, opts ...rerpc.CallOption) (*FailResponse, error) {
	merged := c.mergeOptions(opts)
	call := rerpc.NewClientFunc[FailRequest, FailResponse](
		c.doer,
		c.baseURL,
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"Fail",                 // protobuf method
		merged...,
	)
	return call(ctx, req)
}

// Sum calls internal.ping.v1test.PingService.Sum. Call options passed here
// apply only to this call.
func (c *pingServiceClientReRPC) Sum(ctx context.Context, opts ...rerpc.CallOption) *callstream.Client[SumRequest, SumResponse] {
	merged := c.mergeOptions(opts)
	ctx, call := rerpc.NewClientStream(
		ctx,
		c.doer,
		rerpc.StreamTypeClient,
		c.baseURL,
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"Sum",                  // protobuf method
		merged...,
	)
	stream := call(ctx)
	return callstream.NewClient[SumRequest, SumResponse](stream)
}

// CountUp calls internal.ping.v1test.PingService.CountUp. Call options passed
// here apply only to this call.
func (c *pingServiceClientReRPC) CountUp(ctx context.Context, req *CountUpRequest, opts ...rerpc.CallOption) (*callstream.Server[CountUpResponse], error) {
	merged := c.mergeOptions(opts)
	ctx, call := rerpc.NewClientStream(
		ctx,
		c.doer,
		rerpc.StreamTypeServer,
		c.baseURL,
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"CountUp",              // protobuf method
		merged...,
	)
	stream := call(ctx)
	if err := stream.Send(req); err != nil {
		_ = stream.CloseSend(err)
		_ = stream.CloseReceive()
		return nil, err
	}
	if err := stream.CloseSend(nil); err != nil {
		_ = stream.CloseReceive()
		return nil, err
	}
	return callstream.NewServer[CountUpResponse](stream), nil
}

// CumSum calls internal.ping.v1test.PingService.CumSum. Call options passed
// here apply only to this call.
func (c *pingServiceClientReRPC) CumSum(ctx context.Context, opts ...rerpc.CallOption) *callstream.Bidirectional[CumSumRequest, CumSumResponse] {
	merged := c.mergeOptions(opts)
	ctx, call := rerpc.NewClientStream(
		ctx,
		c.doer,
		rerpc.StreamTypeBidirectional,
		c.baseURL,
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"CumSum",               // protobuf method
		merged...,
	)
	stream := call(ctx)
	return callstream.NewBidirectional[CumSumRequest, CumSumResponse](stream)
}

// PingServiceReRPC is a server for the internal.ping.v1test.PingService
// service. To make sure that adding methods to this protobuf service doesn't
// break all implementations of this interface, all implementations must embed
// UnimplementedPingServiceReRPC.
//
// By default, recent versions of grpc-go have a similar forward compatibility
// requirement. See https://github.com/grpc/grpc-go/issues/3794 for a longer
// discussion.
type PingServiceReRPC interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Fail(context.Context, *FailRequest) (*FailResponse, error)
	Sum(context.Context, *handlerstream.Client[SumRequest, SumResponse]) error
	CountUp(context.Context, *CountUpRequest, *handlerstream.Server[CountUpResponse]) error
	CumSum(context.Context, *handlerstream.Bidirectional[CumSumRequest, CumSumResponse]) error
	mustEmbedUnimplementedPingServiceReRPC()
}

// NewPingServiceHandlerReRPC wraps each method on the service implementation in
// a *rerpc.Handler. The returned slice can be passed to rerpc.NewServeMux.
func NewPingServiceHandlerReRPC(svc PingServiceReRPC, opts ...rerpc.HandlerOption) []*rerpc.Handler {
	handlers := make([]*rerpc.Handler, 0, 5)

	ping := rerpc.NewUnaryHandler(
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"Ping",                 // protobuf method
		svc.Ping,
		opts...,
	)
	handlers = append(handlers, ping)

	fail := rerpc.NewUnaryHandler(
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"Fail",                 // protobuf method
		svc.Fail,
		opts...,
	)
	handlers = append(handlers, fail)

	sum := rerpc.NewStreamingHandler(
		rerpc.StreamTypeClient,
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"Sum",                  // protobuf method
		func(ctx context.Context, sf rerpc.StreamFunc) {
			stream := sf(ctx)
			typed := handlerstream.NewClient[SumRequest, SumResponse](stream)
			err := svc.Sum(stream.Context(), typed)
			_ = stream.CloseReceive()
			if err != nil {
				if _, ok := rerpc.AsError(err); !ok {
					if errors.Is(err, context.Canceled) {
						err = rerpc.Wrap(rerpc.CodeCanceled, err)
					}
					if errors.Is(err, context.DeadlineExceeded) {
						err = rerpc.Wrap(rerpc.CodeDeadlineExceeded, err)
					}
				}
			}
			_ = stream.CloseSend(err)
		},
		opts...,
	)
	handlers = append(handlers, sum)

	countUp := rerpc.NewStreamingHandler(
		rerpc.StreamTypeServer,
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"CountUp",              // protobuf method
		func(ctx context.Context, sf rerpc.StreamFunc) {
			stream := sf(ctx)
			typed := handlerstream.NewServer[CountUpResponse](stream)
			var req CountUpRequest
			if err := stream.Receive(&req); err != nil {
				_ = stream.CloseReceive()
				_ = stream.CloseSend(err)
				return
			}
			if err := stream.CloseReceive(); err != nil {
				_ = stream.CloseSend(err)
				return
			}
			err := svc.CountUp(stream.Context(), &req, typed)
			if err != nil {
				if _, ok := rerpc.AsError(err); !ok {
					if errors.Is(err, context.Canceled) {
						err = rerpc.Wrap(rerpc.CodeCanceled, err)
					}
					if errors.Is(err, context.DeadlineExceeded) {
						err = rerpc.Wrap(rerpc.CodeDeadlineExceeded, err)
					}
				}
			}
			_ = stream.CloseSend(err)
		},
		opts...,
	)
	handlers = append(handlers, countUp)

	cumSum := rerpc.NewStreamingHandler(
		rerpc.StreamTypeBidirectional,
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"CumSum",               // protobuf method
		func(ctx context.Context, sf rerpc.StreamFunc) {
			stream := sf(ctx)
			typed := handlerstream.NewBidirectional[CumSumRequest, CumSumResponse](stream)
			err := svc.CumSum(stream.Context(), typed)
			_ = stream.CloseReceive()
			if err != nil {
				if _, ok := rerpc.AsError(err); !ok {
					if errors.Is(err, context.Canceled) {
						err = rerpc.Wrap(rerpc.CodeCanceled, err)
					}
					if errors.Is(err, context.DeadlineExceeded) {
						err = rerpc.Wrap(rerpc.CodeDeadlineExceeded, err)
					}
				}
			}
			_ = stream.CloseSend(err)
		},
		opts...,
	)
	handlers = append(handlers, cumSum)

	return handlers
}

var _ PingServiceReRPC = (*UnimplementedPingServiceReRPC)(nil) // verify interface implementation

// UnimplementedPingServiceReRPC returns CodeUnimplemented from all methods. To
// maintain forward compatibility, all implementations of PingServiceReRPC must
// embed UnimplementedPingServiceReRPC.
type UnimplementedPingServiceReRPC struct{}

func (UnimplementedPingServiceReRPC) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, rerpc.Errorf(rerpc.CodeUnimplemented, "internal.ping.v1test.PingService.Ping isn't implemented")
}

func (UnimplementedPingServiceReRPC) Fail(context.Context, *FailRequest) (*FailResponse, error) {
	return nil, rerpc.Errorf(rerpc.CodeUnimplemented, "internal.ping.v1test.PingService.Fail isn't implemented")
}

func (UnimplementedPingServiceReRPC) Sum(context.Context, *handlerstream.Client[SumRequest, SumResponse]) error {
	return rerpc.Errorf(rerpc.CodeUnimplemented, "internal.ping.v1test.PingService.Sum isn't implemented")
}

func (UnimplementedPingServiceReRPC) CountUp(context.Context, *CountUpRequest, *handlerstream.Server[CountUpResponse]) error {
	return rerpc.Errorf(rerpc.CodeUnimplemented, "internal.ping.v1test.PingService.CountUp isn't implemented")
}

func (UnimplementedPingServiceReRPC) CumSum(context.Context, *handlerstream.Bidirectional[CumSumRequest, CumSumResponse]) error {
	return rerpc.Errorf(rerpc.CodeUnimplemented, "internal.ping.v1test.PingService.CumSum isn't implemented")
}

func (UnimplementedPingServiceReRPC) mustEmbedUnimplementedPingServiceReRPC() {}
