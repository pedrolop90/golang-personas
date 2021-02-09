package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"primero_test/config"
	"strconv"
	"strings"
)

type Server struct {
	propertiesFactory *config.PropertiesFactory
	controllerFactory *config.ControllerFactory
	serviceFactory    *config.ServiceFactory
	repositoryFactory *config.RepositoryFactory
}

func (s *Server) Init() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	propertiesSystems := s.propertiesFactory.PropertiesSystemFactory()
	personaRepository := s.repositoryFactory.PersonaRepositoryFactory()
	personaService := s.serviceFactory.PersonaServiceFactory(&personaRepository)
	personaController := s.controllerFactory.PersonaControllerFactory(r, &personaService)
	personaController.CreateRouter()

	defaultPort := []string{":", strconv.Itoa(propertiesSystems.Port)}
	http.ListenAndServe(
		strings.Join(
			defaultPort,
			"",
		),
		r,
	)
}
