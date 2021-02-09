package repository

import "primero_test/domain"

type PersonaRepository interface {
	Create(personaDto domain.PersonaDto) domain.PersonaDto

	Update(personaDto domain.PersonaDto) domain.PersonaDto

	FindAll() []domain.PersonaDto

	DeleteById(id int32)
}
