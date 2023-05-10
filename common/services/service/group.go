package service

import (
	"sync"
	"time"
)

type ServiceGroup struct {
	services []Service

	wg *sync.WaitGroup

	e chan error
}

func (s *ServiceGroup) Add(services ...Service) {
	s.services = append(s.services, services...)
}

func (s *ServiceGroup) Run() error {
	s.wg = new(sync.WaitGroup)
	s.wg.Add(len(s.services))
	for _, v := range s.services {
		go s.run(v)
	}

	select {
	case err := <-s.e:
		return err
	case <-time.After(time.Second * 5):
	}
	return nil
}

func (s *ServiceGroup) Shutdown() error {
	for i := len(s.services); i > 0; i-- {
		s.services[i-1].Shutdown()
	}
	s.wg.Wait()
	return nil
}

func (s *ServiceGroup) run(service Service) {
	defer s.wg.Done()
	if err := service.Run(); err != nil {
		s.e <- err
	}
}

func Group(services ...Service) *ServiceGroup {
	return &ServiceGroup{
		services: services,
		e:        make(chan error, len(services)),
	}
}
