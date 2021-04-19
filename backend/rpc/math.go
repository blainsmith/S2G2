package rpc

import (
	"context"

	"github.com/blainsmith/go-svelte-rpc-starter/backend/rpc/rpcserver"
)

type MathService struct{}

func (s *MathService) Add(_ context.Context, r rpcserver.AddRequest) (*rpcserver.AddResponse, error) {
	return nil, nil
}

func (s *MathService) Subtract(_ context.Context, r rpcserver.SubtractRequest) (*rpcserver.SubtractResponse, error) {
	return nil, nil
}
