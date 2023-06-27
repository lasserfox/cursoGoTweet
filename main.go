package main

import (
	"context"
	"cursoGoTweet/awsgo"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/lasserfox/cursoGoTweet/awsgo"
	"github.com/lasserfox/cursoGoTweet/awsgo"
	"os"
)

func main() {
	fmt.Println("Hola")
	lambda.Start(EjecutoLambda)
	awsgo.InicializandoAWS()
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
	SecretModel, err := secretmanager.
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
