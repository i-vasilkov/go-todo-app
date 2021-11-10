package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Uri  string
	User string
	Pass string
}

func NewClient(con Connection) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(con.Uri)
	opts.SetAuth(options.Credential{
		Username: con.User,
		Password: con.Pass,
	})

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
