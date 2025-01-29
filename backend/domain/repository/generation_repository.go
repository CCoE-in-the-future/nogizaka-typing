package repository

import (
	"backend/domain/entities"
)

type GenerationRepository interface {
	GetAll() ([]entities.Generation, error)
	Save(data entities.Generation) (string, error)
}
