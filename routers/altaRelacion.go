package routers

import (
	"context"
	"cursoGoTweet/bd"
	"cursoGoTweet/models"
	"github.com/aws/aws-lambda-go/events"
)

func AltaRelacion(ctx context.Context, request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
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

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		r.Message = "Error al otentar crear la relacion " + err.Error()
		return r
	}
	if !status {
		r.Message = "Error al otentar crear la relacion "
		return r
	}
	r.Status = 200
	r.Message = "Alta relacion OK"
	return r

}
