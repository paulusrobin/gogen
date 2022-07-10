package server

import (
	"github.com/paulusrobin/gogen/internal/config"
	"github.com/paulusrobin/gogen/internal/repository/postgres"
	"github.com/paulusrobin/gogen/internal/repository/postgres/user"
	"os"
)

type httpServer struct {
	sig chan os.Signal
	cfg config.Config
}

func (s httpServer) Run() error {
	db, err := postgres.NewDatabase(s.cfg.Postgres)
	if err != nil {
		return err
	}
	user.NewRepository(db)
	return nil
}

func HTTP(sig chan os.Signal, cfg config.Config) Server {
	return &httpServer{
		sig: sig,
		cfg: cfg,
	}
}
