package datasource

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"invetory-manager-hexagonal/framework/repository"
	"log"

	"os"

	"github.com/joho/godotenv"
)

type ConnectorMongoDataSource struct {
	repository.IConnectorMongoDataSource
}

var clientOptions *options.ClientOptions

func (c *ConnectorMongoDataSource) init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	log.Println("Variáveis foram carregadas")
	url := os.Getenv("MONGO_DATA_BASE_URL")
	log.Printf("MONGO_DATA_BASE_URL=%s", url)
	clientOptions = options.Client().ApplyURI(url)
}

func (c *ConnectorMongoDataSource) Connect() (*mongo.Client, error) {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Connection to MongoDB opened.")
	return client, nil
}

func (c *ConnectorMongoDataSource) DataSource(client *mongo.Client, database, collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}

func (c *ConnectorMongoDataSource) Disconnect(client *mongo.Client) (bool, error) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Println(err)
		return false, err
	}
	log.Println("Connection to MongoDB closed.")
	return true, nil
}

func (c *ConnectorMongoDataSource) Save(collection *mongo.Collection, date interface{}) (string, error) {
	teste, err := collection.InsertOne(context.TODO(), date)
	return teste.InsertedID.(primitive.ObjectID).Hex(), err
}

func NewConnectorMongoDataSource() repository.IConnectorMongoDataSource {
	connectorMongoDataSource := new(ConnectorMongoDataSource)
	connectorMongoDataSource.init()
	return connectorMongoDataSource
}
