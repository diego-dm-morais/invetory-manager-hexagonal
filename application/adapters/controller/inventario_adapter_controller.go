package controller

import (
	"invetory-manager-hexagonal/application/entities"
	"invetory-manager-hexagonal/application/usecase"
)

type InventarioAdapterController struct {
	inventarioApplication usecase.IInventarioApplication
}

func (i InventarioAdapterController) Salvar(inventario InventarioDto) error {
	var itens []entities.Item
	for _, item := range inventario.Itens {
		newItem := entities.NewItem(
			item.Descricao,
			item.Preco,
			item.Quantidade,
		)
		itens = append(itens, newItem)
	}
	newInventario := entities.NewInventario(inventario.Titulo, itens)

	return i.inventarioApplication.Salvar(newInventario, inventario.IdUsuario)
}

func NewInventarioAdapterController(inventarioApplication usecase.IInventarioApplication) *InventarioAdapterController {
	return &InventarioAdapterController{
		inventarioApplication: inventarioApplication,
	}
}
