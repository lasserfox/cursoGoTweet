package routers

import (
	"cursoGoTweet/bd"
	"cursoGoTweet/models"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"strconv"
)

func LeoTweets(request events.APIGatewayProxyRequest) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	pagina := request.QueryStringParameters["pagina"]

	if len(ID) < 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}
	if len(ID) < 1 {
		pagina = "1"
	}
	pag, err := strconv.Atoi(pagina)
	if err != nil {
		r.Message = "Debe enviar en numero en el parámetro pagina"
		return r
	}
	tweets, correcto := bd.LeoTweets(ID, int64(pag))
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
