package usecase

type IPermissaoAdapterApi interface {
	ObterPermissao(tipoPermissao string, idUsuario string) (bool, error)
}
