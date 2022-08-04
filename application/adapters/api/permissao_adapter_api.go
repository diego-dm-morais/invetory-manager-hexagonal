package api

import "invetory-manager-hexagonal/application/usecase"

type PermissaoAdapterApi struct {
	usecase.IPermissaoAdapterApi
}

func (p PermissaoAdapterApi) ObterPermissao(tipoPermissao string, idUsuario string) (bool, error) {
	return true, nil
}

func NewPermissaoAdapterApi() usecase.IPermissaoAdapterApi {
	return PermissaoAdapterApi{}
}
