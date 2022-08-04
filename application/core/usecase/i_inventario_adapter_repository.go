package usecase

import (
	"invetory-manager-hexagonal/application/core/entities"
)

type IInventarioAdapterRepository interface {
	Salvar(inventario entities.Inventario) (string, error)
}
