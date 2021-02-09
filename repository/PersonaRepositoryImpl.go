package repository

import (
	"primero_test/domain"
)

type PersonaRepositoryImpl struct {
	personas []domain.PersonaDto
}

func (p PersonaRepositoryImpl) Create(personaDto domain.PersonaDto) domain.PersonaDto {
	p.personas = append(p.personas, personaDto)
	return personaDto
}

func (p PersonaRepositoryImpl) Update(personaDto domain.PersonaDto) domain.PersonaDto {
	return domain.PersonaDto{}
}

func (p PersonaRepositoryImpl) FindAll() []domain.PersonaDto {
	return nil
}

func (p PersonaRepositoryImpl) DeleteById(id int32) {

}
