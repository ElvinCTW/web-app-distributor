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

type Data struct {
	Id           string  `json:"id"`
	RestaurantId string  `json:"restaurantId"`
	Comment      string  `json:"comment" bson:"comment"`
	Stars        float64 `json:"stars" bson:"stars"`
}

type data struct {
	Oid           primitive.ObjectID `bson:"_id"`
	RestaurantOid primitive.ObjectID `bson:"restaurantOid"`
	*Data
}

func GetDataById(id string) *Data {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}

	ctx := context.Background()
	filter := bson.M{"_id": oid}
	d := new(data)
	if err := col.FindOne(ctx, filter).Decode(d); err != nil && err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		log.Error(err)
		return nil
	} else {
		d.Data.Id = d.Oid.Hex()
		d.Data.RestaurantId = d.RestaurantOid.Hex()
		return d.Data
	}
}
