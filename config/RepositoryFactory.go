package config

import (
	"primero_test/repository"
)

var instanceRepositoryFactory *RepositoryFactory

func NewRepositoryFactory() RepositoryFactory {
	if instanceRepositoryFactory == nil {
		instanceRepositoryFactory = &RepositoryFactory{}
	}
	return *instanceRepositoryFactory
}

type RepositoryFactory struct{}

func (r *RepositoryFactory) PersonaRepositoryFactory() repository.PersonaRepository {
	return repository.PersonaRepositoryImpl{}
}
