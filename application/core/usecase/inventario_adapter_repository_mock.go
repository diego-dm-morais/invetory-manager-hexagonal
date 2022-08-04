package usecase

import (
	"github.com/stretchr/testify/mock"
	"invetory-manager-hexagonal/application/core/entities"
)

type MockInventarioAdapterRepository struct {
	mock.Mock
}

func (mock *MockInventarioAdapterRepository) Salvar(inventario entities.Inventario) (string, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)
}
