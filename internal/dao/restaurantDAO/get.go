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
	d := new(data)
	if err := col.FindOne(ctx, filter).Decode(d); err != nil && err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		log.Error(err)
		return nil
	} else {
		d.Data.Id = d.Oid.Hex()
		return d.Data
	}
}
