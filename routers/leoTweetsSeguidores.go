package routers

import (
	"cursoGoTweet/bd"
	"cursoGoTweet/models"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"strconv"
)

func LeoTweetsSeguidores(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400
	IDUsuario := claim.ID.Hex()

	pagina := request.QueryStringParameters["pagina"]

	if len(pagina) < 1 {
		pagina = "1"
	}
	pag, err := strconv.Atoi(pagina)
	if err != nil {
		r.Message = "Debe enviar en numero en el parámetro pagina"
		return r
	}
	tweets, correcto := bd.LeoTweetsSeguidores(IDUsuario, pag)
	if !correcto {
		r.Message = "Error al leer los tweets."
		return r
	}
	respJson, err := json.Marshal(tweets)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos de los usuarios a JSON"
	}
	r.Status = 200
	r.Message = string(respJson)
	return r
}
