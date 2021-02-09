package config

import (
	"primero_test/repository"
	"primero_test/service"
)

var instanceServiceFactory *ServiceFactory

func NewServiceFactory() ServiceFactory {
	if instanceServiceFactory == nil {
		instanceServiceFactory = &ServiceFactory{}
	}
	return *instanceServiceFactory
}

type ServiceFactory struct {
}

func (s *ServiceFactory) PersonaServiceFactory(personaRepository *repository.PersonaRepository) service.PersonaService {
	return service.PersonaService{PersonaRepository: *personaRepository}
}
