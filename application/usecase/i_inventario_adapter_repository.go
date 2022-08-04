package usecase

import "invetory-manager-hexagonal/application/entities"

type IInventarioAdapterRepository interface {
	Salvar(inventario entities.Inventario) (string, error)
}
