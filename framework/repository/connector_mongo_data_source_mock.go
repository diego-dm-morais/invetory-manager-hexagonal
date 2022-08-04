package repository

import (
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockConnectorMongoDataource struct {
	mock.Mock
}

func (mock *MockConnectorMongoDataource) Connect() (*mongo.Client, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*mongo.Client), args.Error(1)

}
func (mock *MockConnectorMongoDataource) DataSource(client *mongo.Client, database, collection string) *mongo.Collection {
	args := mock.Called()
	result := args.Get(0)
	return result.(*mongo.Collection)
}
func (mock *MockConnectorMongoDataource) Disconnect(client *mongo.Client) (bool, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(bool), args.Error(1)
}

func (mock *MockConnectorMongoDataource) Save(collection *mongo.Collection, date interface{}) (string, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)
}
