package server

import (
	"github.com/paulusrobin/gogen/internal/config"
	"os"
)

type httpServer struct {
	sig chan os.Signal
	cfg config.Config
}

func (s httpServer) Run() {

}

func HTTP(sig chan os.Signal, cfg config.Config) Server {
	return &httpServer{
		sig: sig,
		cfg: cfg,
	}
}
