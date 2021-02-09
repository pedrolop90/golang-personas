package service

import (
	"primero_test/domain"
	"primero_test/repository"
)

type PersonaService struct {
	PersonaRepository repository.PersonaRepository
}

func (p *PersonaService) Created(personaDto domain.PersonaDto) {
	p.PersonaRepository.Create(personaDto)
}
