package database

import (
	"context"
	"fmt"
	"sampleprj/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(c config.AppConfig) (*mongo.Client, error) {
	mongoConnection := fmt.Sprintf("mongodb://%s:%s@%s:%d/", c.MUSER, c.MPASS, c.MHOST, c.MPORT)
	mongo, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoConnection))

	if err != nil {
		return nil, err
	}

	return mongo, nil
}
