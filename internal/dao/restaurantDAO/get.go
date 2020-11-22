package restaurantDAO

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDataById(id string) *Data {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}

	ctx := context.Background()
	filter := bson.M{"_id": oid}
	data := new(Data)
	if err := col.FindOne(ctx, filter).Decode(data); err != nil && err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		log.Error(err)
		return nil
	} else {
		return data
	}
}
