package config

import (
	"github.com/go-chi/chi"
	"primero_test/controller"
	"primero_test/service"
)

var instanceControllerFactory *ControllerFactory

func NewControllerFactory() ControllerFactory {
	if instanceControllerFactory == nil {
		instanceControllerFactory = &ControllerFactory{}
	}
	return *instanceControllerFactory
}

type ControllerFactory struct {
}

func (d ControllerFactory) PersonaControllerFactory(r *chi.Mux, personaService *service.PersonaService) controller.PersonaController {
	return controller.PersonaController{
		Router:         r,
		PersonaService: *personaService,
	}
}
