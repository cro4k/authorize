package service

type Service interface {
	Run() error
	Shutdown() error
}
