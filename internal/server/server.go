package server

type Server interface {
	Run() error
	Shutdown() error
}
