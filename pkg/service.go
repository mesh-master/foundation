package pkg

import (
	"fmt"
	"github.com/go-serv/service/internal/service"
)

type ServiceConfig interface {
	service.ConfigInterface
}

type LocalServiceInterface interface {
	service.LocalServiceInterface
}

type NetworkServiceInterface interface {
	service.NetworkServiceInterface
}

// Service state
type ServiceState service.State

type StateInquirerInterface interface {
	fmt.Stringer
	IsRunning() bool
	IsStopped() bool
	IsSuspended() bool
}

func (s ServiceState) IsRunning() bool {
	return s == ServiceState(service.StateRunning)
}

func (s ServiceState) IsStopped() bool {
	return s == ServiceState(service.StateStopped)
}

func (s ServiceState) IsSuspended() bool {
	return s == ServiceState(service.StateSuspended)
}

func (s ServiceState) String() string {
	return [...]string{"Init", "Running", "StopInProgress", "Stopped", "Suspended"}[s]
}