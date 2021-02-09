package server

import "primero_test/config"

var instanceFabricaServer *ServerFactory

func NewFabricaServer() ServerFactory {
	if instanceFabricaServer == nil {
		instanceFabricaServer = &ServerFactory{}
	}
	return *instanceFabricaServer
}

type ServerFactory struct {
}

func (f ServerFactory) ServerFactory() *Server {

	propertiesFactory := config.NewPropertiesFactory()
	repositoryFactory := config.NewRepositoryFactory()
	serviceFactory := config.NewServiceFactory()
	controllerFactory := config.NewControllerFactory()

	return &Server{
		&propertiesFactory,
		&controllerFactory,
		&serviceFactory,
		&repositoryFactory,
	}
}
