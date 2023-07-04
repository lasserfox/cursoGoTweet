package bd

import (
	"context"
	"fmt"
	"github.com/lasserfox/cursoGoTweet/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN *mongo.Client
var DatabaseName string

func ConectarDB(ctx context.Context) error {
	fmt.Println("Conectado a la DB")
	user := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	fmt.Printf("Url Conexión: mongodb+srv//%s:*****@%s/?retryWrites=true&w=majority", user, host)
	connStr := "mongodb+srv//" + user + ":" + password + "@" + host + "/?retryWrites=true&w=majority"
	fmt.Println("constr: " + connStr)
	fmt.Println("Debería ser: mongodb+srv://user:passw*rd@twitter.oisp1jg.mongodb.net/")
	var clienOptions = options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(ctx, clienOptions)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Conexión exitoda con las BD")
	MongoCN = client
	DatabaseName = ctx.Value(models.Key("database")).(string)
	return nil
}

func BaseConectada() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}
