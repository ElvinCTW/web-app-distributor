package articleDAO

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"web-app-distributor/pkg/logger"
)

var (
	CollectionName = "articleData"
	col            *mongo.Collection
	log            = logger.Get()
)

func Init(db *mongo.Database) {
	col = db.Collection(CollectionName)
}

type Article struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	RestaurantId primitive.ObjectID `json:"restaurantId" bson:"restaurantId"`
	Comment      string
	Stars        float64
}

func Get(id string) *Article {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}

	ctx := context.Background()
	filter := bson.M{"_id": oid}
	data := new(Article)
	if err := col.FindOne(ctx, filter).Decode(data); err != nil && err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		log.Error(err)
		return nil
	} else {
		return data
	}
}
