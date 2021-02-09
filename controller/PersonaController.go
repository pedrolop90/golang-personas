package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"net/http"
	"primero_test/domain"
	"primero_test/service"
)

type PersonaController struct {
	Router         *chi.Mux
	PersonaService service.PersonaService
}

func (p *PersonaController) CreateRouter() {

	p.Router.Route("/personas", func(r chi.Router) {

		r.Get("/", (p).findAll)
		r.Post("/", (p).created)
		r.Put("/{personaId}", (p).update)
		r.Delete("/", (p).delete)

	})

}

func (p *PersonaController) findAll(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola mundo!!!")
}

func (p *PersonaController) created(w http.ResponseWriter, r *http.Request) {
	var persona domain.PersonaDto
	err := json.NewDecoder(r.Body).Decode(&persona)

	if err != nil {
		fmt.Println("fallo", err)
	}
	p.PersonaService.Created(persona)
}

func (p *PersonaController) update(w http.ResponseWriter, r *http.Request) {
	personaId := chi.URLParam(r, "personaId")
	io.WriteString(w, personaId)

	var persona domain.PersonaDto
	err := json.NewDecoder(r.Body).Decode(&persona)

	if err != nil {
		fmt.Println("fallo", err)
	}
	fmt.Println(persona)
}

func (p *PersonaController) delete(w http.ResponseWriter, r *http.Request) {
	personaId := chi.URLParam(r, "personaId")
	io.WriteString(w, personaId)
}
