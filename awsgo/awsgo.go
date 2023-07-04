package awsgo

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func InicializandoAWS() {
	fmt.Println("Iniciando AWS, creando context")
	Ctx = context.TODO()
	fmt.Println("Context creado, load Config")
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))
	fmt.Println("Config : " + Cfg.Region)
	if err != nil {
		panic("Error al cargar la configuración de .aws/config " + err.Error())
	}
}
