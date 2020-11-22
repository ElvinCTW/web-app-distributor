package restaurantDAO

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"web-app-distributor/pkg/logger"
)

var (
	CollectionName = "restaurantData"
	col            *mongo.Collection
	log            = logger.Get()
)

func Init(db *mongo.Database) {
	col = db.Collection(CollectionName)
}

type Data struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Info      Info               `json:"info"`
	Dishes    []Dish
	PictureId []string
}

type Info struct {
	Address  string
	City     string
	Distinct string
	Genre    string
}

type Dish struct {
	Name      string
	Price     string
	PictureId []string
}
