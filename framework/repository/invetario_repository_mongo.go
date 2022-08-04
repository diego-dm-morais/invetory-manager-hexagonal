package repository

import (
	"errors"
	"invetory-manager-hexagonal/application/adapters/repository"
	"invetory-manager-hexagonal/application/core/entities"
)

type InvetarioRepositoryMongo struct {
	repository.IInvetarioRepositoryMongo
	connectorDataSource IConnectorMongoDataSource
}

const DATA_BASE_LABSIT, TABLE_ORDERS = "labsit", "inventory"

func (i InvetarioRepositoryMongo) Inserir(inventario entities.Inventario) (string, error) {
	client, err := i.connectorDataSource.Connect()

	if err != nil {
		return "", errors.New("conexão com banco de dados falhou")
	}
	defer i.connectorDataSource.Disconnect(client)
	collection := i.connectorDataSource.DataSource(client, DATA_BASE_LABSIT, TABLE_ORDERS)
	id, err := i.connectorDataSource.Save(collection, inventario)
	return id, err
}

func NewInvetarioRepositoryMongo(connectorDataSource IConnectorMongoDataSource) repository.IInvetarioRepositoryMongo {
	return &InvetarioRepositoryMongo{connectorDataSource: connectorDataSource}
}
