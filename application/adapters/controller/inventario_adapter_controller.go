package controller

import (
	entities2 "invetory-manager-hexagonal/application/core/entities"
	"invetory-manager-hexagonal/application/core/usecase"
)

type InventarioAdapterController struct {
	inventarioApplication usecase.IInventarioApplication
}

func (i InventarioAdapterController) Salvar(inventario InventarioDto) error {
	var itens []entities2.Item
	for _, item := range inventario.Itens {
		newItem := entities2.NewItem(
			item.Descricao,
			item.Preco,
			item.Quantidade,
		)
		itens = append(itens, newItem)
	}
	newInventario := entities2.NewInventario(inventario.Titulo, itens)

	return i.inventarioApplication.Salvar(newInventario, inventario.IdUsuario)
}

func NewInventarioAdapterController(inventarioApplication usecase.IInventarioApplication) *InventarioAdapterController {
	return &InventarioAdapterController{
		inventarioApplication: inventarioApplication,
	}
}
