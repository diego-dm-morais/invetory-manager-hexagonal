package usecase

import "invetory-manager-hexagonal/application/entities"

type InventarioUseCase struct {
	inventarioAdapter IInventarioAdapterRepository
	permissaoAdapter  IPermissaoAdapterApi
}

func (i InventarioUseCase) Salvar(inventario entities.Inventario, idUsuario string) error {
	isValid, err := inventario.EValido()
	if isValid == false {
		return err
	}

	allowed, err := i.permissaoAdapter.ObterPermissao("CRIAR_INVENTARIO", idUsuario)
	if allowed == false {
		return err
	}

	id, err := i.inventarioAdapter.Salvar(inventario)
	if len(id) == 0 {
		return err
	}

	return nil
}

func NewInventarioApplication(datasource IInventarioAdapterRepository, permissaoAdapter IPermissaoAdapterApi) IInventarioApplication {
	return &InventarioUseCase{
		inventarioAdapter: datasource,
		permissaoAdapter:  permissaoAdapter,
	}
}
