package routers

import (
	"cursoGoTweet/bd"
	"cursoGoTweet/models"
	"github.com/aws/aws-lambda-go/events"
)

func EliminarTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]

	if len(ID) < 1 {
		r.Message = "ID obligarotio"
		return r
	}

	err := bd.BorrarTweet(ID, claim.ID.Hex())
	if err != nil {
		r.Message = "Error al borrar el tweet " + err.Error()
		return r
	}
	r.Status = 200
	r.Message = "Tweet Eliminado correctamente"
	return r

}
