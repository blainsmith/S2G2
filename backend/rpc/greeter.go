package rpc

import (
	"context"

	"github.com/blainsmith/go-svelte-rpc-starter/backend/rpc/rpcserver"
)

type GreeterService struct{}

func (s *GreeterService) Greet(_ context.Context, _ rpcserver.GreetRequest) (*rpcserver.GreetResponse, error) {
	return nil, nil
}
