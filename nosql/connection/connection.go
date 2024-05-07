package connection

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

type Config struct {
	User, Host, Psw, Name, Port string
	GormConfig                  gorm.Config
}

type Connection struct {
	*mongo.Client
}

func (c *Connection) Connect(ctx context.Context, host, port string) (err error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s/", host, port))
	c.Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return
	}
	err = c.Ping(ctx, nil)
	if err != nil {
		return
	}
	return
}
func (c *Connection) GetDatabaseCollection(database, collection string) *mongo.Collection {
	return c.Client.Database(database).Collection(collection)
}
