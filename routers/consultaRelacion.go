package routers

import (
	"cursoGoTweet/bd"
	"cursoGoTweet/models"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func ConsultaRelacion(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "ID obligarotio"
		return r
	}

	var t models.Relacion
	t.UsuarioID = claim.ID.Hex()
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion
	resp.Status = bd.ConsultoRelacion(t)

	respJson, err := json.Marshal(resp.Status)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos de los usuarios como json " + err.Error()
		return r
	}
	r.Status = 200
	r.Message = string(respJson)
	return r
}
