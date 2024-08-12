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

func StartTransaction(ctx context.Context) (mongo.Session, error) {
	var session mongo.Session
	var err error
	if session, err = conn.client.StartSession(); err != nil {
		return nil, err
	}
	if err = session.StartTransaction(); err != nil {
		return nil, err
	}
	return session, nil
}

func CommitTransaction(ctx context.Context, session mongo.Session) {
	session.CommitTransaction(ctx)
}

func RollbackTransaction(ctx context.Context, session mongo.Session) {
	session.AbortTransaction(ctx)
}
