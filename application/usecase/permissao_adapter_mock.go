package usecase

import (
	"github.com/stretchr/testify/mock"
)

type MockPermissaoAdapter struct {
	mock.Mock
}

func (mock *MockPermissaoAdapter) ObterPermissao(tipoPermissao string, idUsuario string) (bool, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(bool), args.Error(1)
}
