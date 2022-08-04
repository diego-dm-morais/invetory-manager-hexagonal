package repository

import (
	"invetory-manager-hexagonal/application/core/entities"
)

type IInvetarioRepositoryMongo interface {
	Inserir(inventario entities.Inventario) (string, error)
}
