package repository

import (
	"github.com/google/uuid"

	"backend/domain/entities"
	"backend/domain/repository"
)

type MockGenerationRepository struct {
	mockData map[string]entities.Generation
}

// NewMockGenerationRepository creates a new instance of MockGenerationRepository
func NewMockGenerationRepository() repository.GenerationRepository {
	return &MockGenerationRepository{
		mockData: make(map[string]entities.Generation),
	}
}

// GetAll retrieves all mock data from the repository
func (m *MockGenerationRepository) GetAll() ([]entities.Generation, error) {
	results := make([]entities.Generation, 0, len(m.mockData)) // Create a slice for results
	for _, value := range m.mockData {
		results = append(results, value)
	}
	return results, nil
}

// Save stores data in the repository with a generated UUID as the key
func (m *MockGenerationRepository) Save(data entities.Generation) (string, error) {
	id := uuid.New().String() // Generate a UUID
	m.mockData[id] = data     // Save data with UUID as the key
	return id, nil            // Return the generated UUID
}

// GetByID retrieves data by its ID from the repository
// func (m *MockGenerationRepository) GetByID(id string) (entities.Generation, error) {
// 	if data, exists := m.mockData[id]; exists {
// 		return data, nil
// 	}
// 	return nil, fmt.Errorf("data with ID %s not found", id)
// }
