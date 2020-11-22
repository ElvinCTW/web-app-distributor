package restaurantDAO

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(d *Data) string {
	ctx := context.Background()
	if result, err := col.InsertOne(ctx, d); err != nil {
		log.Error(err)
		return ""
	} else {
		return result.InsertedID.(primitive.ObjectID).Hex()
	}
}
