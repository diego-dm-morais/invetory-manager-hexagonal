package repository

import (
	"invetory-manager-hexagonal/application/entities"
	"invetory-manager-hexagonal/application/usecase"
)

type InventarioAdapterRepository struct {
	inventarioRepositoryMongo IInvetarioRepositoryMongo
}

func (a InventarioAdapterRepository) Salvar(inventario entities.Inventario) (string, error) {
	return a.inventarioRepositoryMongo.Inserir(inventario)
}

func NewInventarioAdapterRepository(inventarioRepository IInvetarioRepositoryMongo) usecase.IInventarioAdapterRepository {
	return &InventarioAdapterRepository{inventarioRepositoryMongo: inventarioRepository}
}
