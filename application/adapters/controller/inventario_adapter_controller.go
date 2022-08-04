package controller

import (
	core_entities "invetory-manager-hexagonal/application/core/entities"
	"invetory-manager-hexagonal/application/core/usecase"
)

type InventarioAdapterController struct {
	inventarioApplication usecase.IInventarioApplication
}

func (i InventarioAdapterController) Salvar(inventario InventarioDto) error {
	var itens []core_entities.Item
	for _, item := range inventario.Itens {
		newItem := core_entities.NewItem(
			item.Descricao,
			item.Preco,
			item.Quantidade,
		)
		itens = append(itens, newItem)
	}
	newInventario := core_entities.NewInventario(inventario.Titulo, itens)

	return i.inventarioApplication.Salvar(newInventario, inventario.IdUsuario)
}

func NewInventarioAdapterController(inventarioApplication usecase.IInventarioApplication) *InventarioAdapterController {
	return &InventarioAdapterController{
		inventarioApplication: inventarioApplication,
	}
}
