package repository

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
	inventarioRepository := new(MockInvetarioRepositoryMongo)
	inventarioAdapter := NewInventarioAdapterRepository(inventarioRepository)

	inventarioRepository.On("Inserir").Return("05919a92-e4df-406f-825a-1da683182ed9", nil)

	itens := []entities2.Item{
		entities2.NewItem("Chocolate", 10.3, 20),
		entities2.NewItem("Refrigerante", 7.4, 100),
		entities2.NewItem("Bolacha", 4.5, 70),
	}

	inventario := entities2.NewInventario("Inventário 2022", itens)

	id, err := inventarioAdapter.Salvar(inventario)

	assert.Equal("05919a92-e4df-406f-825a-1da683182ed9", id)
	assert.NoError(err)
	mock.AssertExpectationsForObjects(t, inventarioRepository)
	inventarioRepository.AssertNumberOfCalls(t, "Inserir", 1)
}

func Test_devera_simular_um_erro_ao_salvar(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	inventarioRepository := new(MockInvetarioRepositoryMongo)
	inventarioAdapter := NewInventarioAdapterRepository(inventarioRepository)

	inventarioRepository.On("Inserir").Return("", errors.New("erro ocorrido ao tentar salvar o inventário"))

	itens := []entities2.Item{
		entities2.NewItem("Chocolate", 10.3, 20),
		entities2.NewItem("Refrigerante", 7.4, 100),
		entities2.NewItem("Bolacha", 4.5, 70),
	}

	inventario := entities2.NewInventario("Inventário 2022", itens)

	id, err := inventarioAdapter.Salvar(inventario)

	assert.Empty(id)
	mock.AssertExpectationsForObjects(t, inventarioRepository)
	assert.Errorf(err, "erro ocorrido ao tentar salvar o inventário")
	inventarioRepository.AssertNumberOfCalls(t, "Inserir", 1)
}
