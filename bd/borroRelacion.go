package bd

import (
	"context"
	"cursoGoTweet/models"
)

func BorroRelacion(t models.Relacion) (bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		{
			return false, err
		}
	}

	return true, nil
}
