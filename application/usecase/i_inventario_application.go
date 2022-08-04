package usecase

import "invetory-manager-hexagonal/application/entities"

type IInventarioApplication interface {
	Salvar(inventario entities.Inventario, usuario string) error
}
