package net

import (
	"context"
	"github.com/mesh-master/foundation/internal/grpc/msg/response"
	"github.com/mesh-master/foundation/internal/runtime"
	"github.com/mesh-master/foundation/pkg/z"
	"github.com/mesh-master/foundation/pkg/z/event"
	"google.golang.org/grpc"
)

func NewServerContext(ctx context.Context, svc z.NetworkServiceInterface, req z.RequestInterface, handler grpc.UnaryHandler) *srvContext {
	srvCtx := new(srvContext)
	srvCtx.Context = ctx
	srvCtx.service = svc
	srvCtx.req = req
	res, _ := response.NewResponse(nil, nil)
	srvCtx.res = res
	srvCtx.handler = handler
	return srvCtx
}

func NewClientContext(ctx context.Context) *clientCtx {
	c := new(clientCtx)
	c.Context = ctx
	runtime.Runtime().TriggerEvent(event.NetClientNewContext, c)
	return c
}
