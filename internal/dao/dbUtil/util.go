package dbUtil

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"web-app-distributor/pkg/logger"
)

var log = logger.Get()

func GetCore(id string, data interface{}, col *mongo.Collection) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(err)
		return
	}

	ctx := context.Background()
	filter := bson.M{"_id": oid}
	if err := col.FindOne(ctx, filter).Decode(data); err != nil && err == mongo.ErrNoDocuments {
		return
	} else if err != nil {
		log.Error(err)
		return
	}
}
