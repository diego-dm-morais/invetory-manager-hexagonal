package repository

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	core_entities "invetory-manager-hexagonal/application/core/entities"
	"testing"
)

func Test_salvar_inventario_com_conector_mongobd(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	connectorMongoDataource := new(MockConnectorMongoDataource)
	invetarioRepositoryMongo := NewInvetarioRepositoryMongo(connectorMongoDataource)

	connectorMongoDataource.On("Connect").Return(&mongo.Client{}, nil)
	connectorMongoDataource.On("DataSource").Return(&mongo.Collection{})
	connectorMongoDataource.On("Disconnect").Return(true, nil)
	connectorMongoDataource.On("Save").Return("05919a92-e4df-406f-825a-1da683182ed9", nil)

	itens := []core_entities.Item{
		core_entities.NewItem("Chocolate", 10.3, 20),
		core_entities.NewItem("Refrigerante", 7.4, 100),
		core_entities.NewItem("Bolacha", 4.5, 70),
	}

	inventario := core_entities.NewInventario("Inventário 2022", itens)

	id, err := invetarioRepositoryMongo.Inserir(inventario)

	assert.Equal("05919a92-e4df-406f-825a-1da683182ed9", id)
	assert.NoError(err)

	mock.AssertExpectationsForObjects(t, connectorMongoDataource)
	connectorMongoDataource.AssertNumberOfCalls(t, "Connect", 1)
	connectorMongoDataource.AssertNumberOfCalls(t, "DataSource", 1)
	connectorMongoDataource.AssertNumberOfCalls(t, "Disconnect", 1)
	connectorMongoDataource.AssertNumberOfCalls(t, "Save", 1)

}

func Test_salvar_inventario_com_error(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	connectorMongoDataource := new(MockConnectorMongoDataource)
	invetarioRepositoryMongo := NewInvetarioRepositoryMongo(connectorMongoDataource)

	connectorMongoDataource.On("Connect").Return(&mongo.Client{}, nil)
	connectorMongoDataource.On("DataSource").Return(&mongo.Collection{})
	connectorMongoDataource.On("Disconnect").Return(true, nil)
	connectorMongoDataource.On("Save").Return("", errors.New("erro ao persistir dados na base"))

	itens := []core_entities.Item{
		core_entities.NewItem("Chocolate", 10.3, 20),
		core_entities.NewItem("Refrigerante", 7.4, 100),
		core_entities.NewItem("Bolacha", 4.5, 70),
	}

	inventario := core_entities.NewInventario("Inventário 2022", itens)

	id, err := invetarioRepositoryMongo.Inserir(inventario)
	assert.Errorf(err, err.Error())
	assert.Empty(id)

	mock.AssertExpectationsForObjects(t, connectorMongoDataource)
	connectorMongoDataource.AssertNumberOfCalls(t, "Connect", 1)
	connectorMongoDataource.AssertNumberOfCalls(t, "DataSource", 1)
	connectorMongoDataource.AssertNumberOfCalls(t, "Disconnect", 1)
	connectorMongoDataource.AssertNumberOfCalls(t, "Save", 1)

}
