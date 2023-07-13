package bd

import (
	"context"
	"cursoGoTweet/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuarios")

	var results []*models.Usuario

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSkip((page - 1) * 20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, opciones)

	if err != nil {
		return results, false
	}

	var incluir bool

	for cursor.Next(ctx) {
		var s models.Usuario
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println("Error en Decode : " + err.Error())
			return results, false
		}
		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false
		encontrado := ConsultoRelacion(r)
		if tipo == "new" && !encontrado {
			incluir = true
		}
		if tipo == "follow" && encontrado {
			incluir = true
		}
		if r.UsuarioRelacionID == ID {
			incluir = false
		}
		if incluir {
			s.Password = ""
			results = append(results, &s)
		}
	}
	err = cursor.Err()
	if err != nil {
		fmt.Println("Error en el cursos " + err.Error())
		return results, false
	}
	cursor.Close(ctx)
	return results, true
}
