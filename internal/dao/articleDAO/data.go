package articleDAO

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"web-app-distributor/internal/dao/dbUtil"
)

var (
	CollectionName = "articleData"
	col            *mongo.Collection
)

func Init(collection *mongo.Collection) {
	col = collection
}

type Article struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	RestaurantId primitive.ObjectID `json:"restaurantId" bson:"restaurantId"`
	Comment      string
	Stars        float64
}

func Get(id string) *Article {
	resp := new(Article)
	dbUtil.GetCore(id, resp, col)
	return resp
}
