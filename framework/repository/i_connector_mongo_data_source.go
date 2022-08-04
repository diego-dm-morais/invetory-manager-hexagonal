package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type IConnectorMongoDataSource interface {
	Connect() (*mongo.Client, error)
	DataSource(client *mongo.Client, database, collection string) *mongo.Collection
	Disconnect(client *mongo.Client) (bool, error)
	Save(collection *mongo.Collection, date interface{}) (string, error)
}
