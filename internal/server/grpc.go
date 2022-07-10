package server

import (
	"github.com/paulusrobin/gogen/internal/config"
	"os"
)

type grpcServer struct {
	sig chan os.Signal
	cfg config.Config
}

func (s grpcServer) Run() error {
	panic("implement me")
}

func GRPC(sig chan os.Signal, cfg config.Config) Server {
	return &grpcServer{
		sig: sig,
		cfg: cfg,
	}
}
