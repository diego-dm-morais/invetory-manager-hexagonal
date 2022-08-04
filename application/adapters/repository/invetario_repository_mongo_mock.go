package repository

import (
	"github.com/stretchr/testify/mock"
	"invetory-manager-hexagonal/application/core/entities"
)

type MockInvetarioRepositoryMongo struct {
	mock.Mock
}

func (mock *MockInvetarioRepositoryMongo) Inserir(inventario entities.Inventario) (string, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)
}
