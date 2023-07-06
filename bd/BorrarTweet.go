package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorrarTweet(ID string, UserID string) error {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id":    objID,
		"userid": UserID,
	}
	_, err := col.DeleteOne(ctx, condicion)
	if err != nil {
		return err
	}
	return nil

}
