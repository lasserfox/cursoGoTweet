package routers

import (
	"cursoGoTweet/bd"
	"cursoGoTweet/models"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"strconv"
)

func ListaUsuarios(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	page := request.QueryStringParameters["page"]
	typeUser := request.QueryStringParameters["type"]
	search := request.QueryStringParameters["search"]
	IDUsuario := claim.ID.Hex()

	if len(page) == 0 {
		page = "1"
	}

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "Debe enviar el el parametro page mayor a o " + err.Error()
		return r
	}

	usuarios, status := bd.LeoUsuariosTodos(IDUsuario, int64(pageTemp), search, typeUser)
	if !status {
		r.Message = "Error al leer los datos de los usuarios"
		return r
	}

	respJson, err := json.Marshal(usuarios)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos a JSON" + err.Error()
		return r
	}
	r.Status = 200
	r.Message = string(respJson)
	return r
}
