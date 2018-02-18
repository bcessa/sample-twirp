package rpc

import (
	"context"
	"github.com/bcessa/sample-twirp/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

type SampleServer struct{}

func (s *SampleServer) Ping(context.Context, *empty.Empty) (*sample.Pong, error) {
	return &sample.Pong{Ok: true}, nil
}
