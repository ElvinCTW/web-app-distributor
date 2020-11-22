package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"web-app-distributor/config"
	"web-app-distributor/internal/dao/articleDAO"
	"web-app-distributor/internal/dao/restaurantDAO"
)

var c = config.Get()

func Init() {
	if client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.DatabaseURL)); err != nil {
		panic(err.Error())
	} else {
		db := client.Database(c.Database)
		articleDAO.Init(db)
		restaurantDAO.Init(db)
	}
}
