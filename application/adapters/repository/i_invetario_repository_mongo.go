package repository

import "invetory-manager-hexagonal/application/entities"

type IInvetarioRepositoryMongo interface {
	Inserir(inventario entities.Inventario) (string, error)
}
