package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConnection struct {
	client *mongo.Client
}

var conn *MongoDBConnection

func InitConnection() error {
	if conn == nil {
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		user := "admin"
		pass := "admin"
		uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.v89st.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", user, pass)
		opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
		// Create a new client and connect to the server
		client, err := mongo.Connect(context.TODO(), opts)
		if err != nil {
			panic(err)
		}
		conn = &MongoDBConnection{
			client: client,
		}
	}
	return nil
}

func GetClient() *MongoDBConnection {
	return conn
}

func (c MongoDBConnection) GetCollection(name string) *mongo.Collection {
	return c.client.Database("minitweeter").Collection(name)
}
