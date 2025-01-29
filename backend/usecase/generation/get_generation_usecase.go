package usecase

import (
	"backend/domain/entities"
	"backend/domain/repository"
)

type GenerationUseCase struct {
	generationRepository repository.GenerationRepository
}

func NewGenerationUseCase(gri repository.GenerationRepository) *GenerationUseCase {
	return &GenerationUseCase{generationRepository: gri}
}

func (guc *GenerationUseCase) GetAllGenerations() ([]entities.Generation, error) {
	guc.generationRepository.Save(entities.Generation{Generation: "6th", Image: "Unknown"})
	generations, err := guc.generationRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return generations, err
}
