// Code generated by goctl. DO NOT EDIT.
// Source: zrpc.proto

package zrpcservice

import (
	"context"

	"eventos-tec/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Event              = __.Event
	ListEventsRequest  = __.ListEventsRequest
	ListEventsResponse = __.ListEventsResponse

	ZrpcService interface {
		ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error)
	}

	defaultZrpcService struct {
		cli zrpc.Client
	}
)

func NewZrpcService(cli zrpc.Client) ZrpcService {
	return &defaultZrpcService{
		cli: cli,
	}
}

func (m *defaultZrpcService) ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error) {
	client := __.NewZrpcServiceClient(m.cli.Conn())
	return client.ListEvents(ctx, in, opts...)
}