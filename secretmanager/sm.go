package secretmanager

// Use this code snippet in your app.
// If you need more information about configurations or implementing the sample code, visit the AWS docs:
// https://aws.github.io/aws-sdk-go-v2/docs/getting-started/

import (
	"cursoGoTweet/awsgo"
	"cursoGoTweet/models"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var datosSecret models.Secret
	fmt.Println("> Pidiendo Secrets " + secretName)
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	fmt.Println("> Config svc creado")
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println("error al obtener el secret " + err.Error())
		return datosSecret, err
	}
	fmt.Println("secret conseguida, haciendo unmarshall ")
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("Lectura de secret OK " + secretName)
	return datosSecret, nil
}
