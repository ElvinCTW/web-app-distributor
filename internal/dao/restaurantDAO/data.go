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
	Id       string   `json:"id" bson:"id"`
	Info     Info     `json:"info" bson:"info"`
	Dishes   []Dish   `json:"dishes" bson:"dishes"`
	Pictures []string `json:"pictures" bson:"pictures"`
}

type data struct {
	Oid primitive.ObjectID `bson:"_id"`
	*Data
}

type Info struct {
	Address  string `json:"address" bson:"address"`
	City     string `json:"city" bson:"city"`
	Distinct string `json:"distinct" bson:"distinct"`
	Genre    string `json:"genre" bson:"genre"`
}

type Dish struct {
	Name     string   `json:"name" bson:"name"`
	Price    int      `json:"price" bson:"price"`
	Pictures []string `json:"pictures" bson:"pictures"`
}
