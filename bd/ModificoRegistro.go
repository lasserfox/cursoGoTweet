package bd

import (
	"context"
	"cursoGoTweet/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuarios")
	fmt.Println("Entro modifico Registro")
	registro := make(map[string]interface{})

	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}
	registro["fechaNacimiento"] = u.FechaNacimiento

	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioweb"] = u.SitioWeb
	}
	fmt.Println("preparando updString ")
	updtString := bson.M{
		"$set": registro,
	}

	fmt.Println("preparando objID ")
	objID, _ := primitive.ObjectIDFromHex(ID)
	fmt.Println("preparando filtro ")
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	bsonData, err := bson.Marshal(filtro)
	if err != nil {
		fmt.Println("Error al convertir a BSON:", err)
	}
	strData := string(bsonData)
	fmt.Println("filtro: " + strData)

	_, err = col.UpdateOne(ctx, filtro, updtString)

	if err != nil {
		return false, err
	}
	return true, nil
}
