package rpc

import (
	"context"
	"github.com/bcessa/sample-twirp/proto"
)

type SampleServer struct{}

func (s *SampleServer) Ping(context.Context, *sample.Empty) (*sample.Pong, error) {
	return &sample.Pong{Ok: true}, nil
}
