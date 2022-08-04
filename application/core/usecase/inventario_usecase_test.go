package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	entities2 "invetory-manager-hexagonal/application/core/entities"
	"testing"
)

func Test_devera_salvar_inventario_com_sucesso(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	permissaoAdapter := new(MockPermissaoAdapter)
	invetarioDatasource := new(MockInventarioAdapterRepository)

	permissaoAdapter.On("ObterPermissao").Return(true, nil)
	invetarioDatasource.On("Salvar").Return("d1653b20-cf96-4134-9e0d-8cac9b37cac3", nil)

	var invetarioApplication IInventarioApplication = NewInventarioApplication(invetarioDatasource, permissaoAdapter)

	itens := []entities2.Item{
		entities2.NewItem("Chocolate", 10.3, 20),
		entities2.NewItem("Refrigerante", 7.4, 100),
		entities2.NewItem("Bolacha", 4.5, 70),
	}

	inventario := entities2.NewInventario("Inventário 2022", itens)
	err := invetarioApplication.Salvar(inventario, "f5b21367-226e-4df9-b36b-bb4dd4bdddb1")

	mock.AssertExpectationsForObjects(t, permissaoAdapter, invetarioDatasource)
	assert.NoError(err)
	permissaoAdapter.AssertNumberOfCalls(t, "ObterPermissao", 1)
	invetarioDatasource.AssertNumberOfCalls(t, "Salvar", 1)

}

func Test_devera_salvar_inventario_com_item_com_preco_igual_zero(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	permissaoAdapter := new(MockPermissaoAdapter)
	invetarioDatasource := new(MockInventarioAdapterRepository)

	var invetarioApplication IInventarioApplication = NewInventarioApplication(invetarioDatasource, permissaoAdapter)

	itens := []entities2.Item{
		entities2.NewItem("Chocolate", 0, 20),
		entities2.NewItem("Refrigerante", 7.4, 100),
		entities2.NewItem("Bolacha", 4.5, 70),
	}

	inventario := entities2.NewInventario("Inventário 2022", itens)
	err := invetarioApplication.Salvar(inventario, "f5b21367-226e-4df9-b36b-bb4dd4bdddb1")

	mock.AssertExpectationsForObjects(t, permissaoAdapter, invetarioDatasource)
	assert.Errorf(err, "preço do item é inválido")
	permissaoAdapter.AssertNotCalled(t, "ObterPermissao")
	invetarioDatasource.AssertNotCalled(t, "Salvar")

}

func Test_devera_tentar_salvar_com_um_usuario_sem_permissao(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	permissaoAdapter := new(MockPermissaoAdapter)
	invetarioDatasource := new(MockInventarioAdapterRepository)

	permissaoAdapter.On("ObterPermissao").Return(false, errors.New("usuário sem permissão"))

	var invetarioApplication IInventarioApplication = NewInventarioApplication(invetarioDatasource, permissaoAdapter)

	itens := []entities2.Item{
		entities2.NewItem("Chocolate", 10.3, 20),
		entities2.NewItem("Refrigerante", 7.4, 100),
		entities2.NewItem("Bolacha", 4.5, 70),
	}

	inventario := entities2.NewInventario("Inventário 2022", itens)
	err := invetarioApplication.Salvar(inventario, "f5b21367-226e-4df9-b36b-bb4dd4bdddb1")

	mock.AssertExpectationsForObjects(t, permissaoAdapter, invetarioDatasource)
	assert.Errorf(err, "usuário sem permissão")
	permissaoAdapter.AssertNumberOfCalls(t, "ObterPermissao", 1)
	invetarioDatasource.AssertNotCalled(t, "Salvar")

}

func Test_devera_tratar_o_erro_retornado_pelo_datasource(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	permissaoAdapter := new(MockPermissaoAdapter)
	invetarioDatasource := new(MockInventarioAdapterRepository)

	permissaoAdapter.On("ObterPermissao").Return(true, nil)
	invetarioDatasource.On("Salvar").Return("", errors.New("erro ocorrido ao tentar salvar inventário"))

	var invetarioApplication IInventarioApplication = NewInventarioApplication(invetarioDatasource, permissaoAdapter)

	itens := []entities2.Item{
		entities2.NewItem("Chocolate", 10.3, 20),
		entities2.NewItem("Refrigerante", 7.4, 100),
		entities2.NewItem("Bolacha", 4.5, 70),
	}

	inventario := entities2.NewInventario("Inventário 2022", itens)
	err := invetarioApplication.Salvar(inventario, "f5b21367-226e-4df9-b36b-bb4dd4bdddb1")

	mock.AssertExpectationsForObjects(t, permissaoAdapter, invetarioDatasource)
	assert.Errorf(err, "erro ocorrido ao tentar salvar inventário")
	permissaoAdapter.AssertNumberOfCalls(t, "ObterPermissao", 1)
	invetarioDatasource.AssertNumberOfCalls(t, "Salvar", 1)

}
