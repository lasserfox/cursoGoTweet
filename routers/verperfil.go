package routers

import (
	"cursoGoTweet/bd"
	"cursoGoTweet/models"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func VerPerfil(request events.APIGatewayProxyRequest) models.RespApi {
	var r models.RespApi
	r.Status = 400

	fmt.Println("Entrando a VerPerfil")
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		r.Message = "Error al buscar id en la BD " + err.Error()
		return r
	}
	respJson, err := json.Marshal(perfil)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos de los usuarios como JSON " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
