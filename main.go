package main

import (
	"context"
	"cursoGoTweet/bd"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lasserfox/cursoGoTweet/awsgo"
	"github.com/lasserfox/cursoGoTweet/handlers"
	"github.com/lasserfox/cursoGoTweet/models"
	"github.com/lasserfox/cursoGoTweet/secretmanager"
	"os"
	"strings"
)

func main() {
	fmt.Println("Entrando al main")
	awsgo.InicializandoAWS()
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse
	if !ValidoParametros() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno. Deben incluir 'SecretName', 'BucketName' y 'UrlPrefix'",
			Headers: map[string]string{
				"Context-Type": "application/json",
			},
		}
		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura del secret 'SecretName'" + err.Error(),
			Headers: map[string]string{
				"Context-Type": "application/json",
			},
		}
		return res, nil
	}
	// TODO: Mirar  si es twittergo o twitterGo
	path := strings.Replace(request.PathParameters["twittergo"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	// Chequeo de conexión a la BD o conecto la BD
	err = bd.ConectarDB(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error en la conexión a la BD" + err.Error(),
			Headers: map[string]string{
				"Context-Type": "application/json",
			},
		}
		return res, nil
	}
	respAPI := handlers.Manejadores(awsgo.Ctx, request)
	if respAPI.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: respAPI.Status,
			Body:       respAPI.Message,
			Headers: map[string]string{
				"Context-Type": "application/json",
			},
		}
		return res, nil
	} else {
		return respAPI.CustomResp, nil
	}

}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return false
	}
	_, traeParametro = os.LookupEnv("BucketName")
	if !traeParametro {
		return false
	}
	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return false
	}
	return true
}
