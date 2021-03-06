package server

import (
	"github.com/paulusrobin/gogen/internal/config"
	"os"
)

type subscriberServer struct {
	sig chan os.Signal
	cfg config.Config
}

func (s *subscriberServer) Run() error {
	panic("implement me")
}

func (s *subscriberServer) Shutdown() error {
	panic("implement me")
}

func Subscriber(sig chan os.Signal, cfg config.Config) (Server, error) {
	return &subscriberServer{
		sig: sig,
		cfg: cfg,
	}, nil
}
