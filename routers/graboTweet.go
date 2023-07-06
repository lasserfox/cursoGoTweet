package routers

import (
	"context"
	"cursoGoTweet/bd"
	"cursoGoTweet/models"
	"encoding/json"
	"time"
)

func GraboTweet(ctx context.Context, claim models.Claim) models.RespApi {
	var mensaje models.Tweet
	var r models.RespApi
	r.Status = 400

	IDUsuario := claim.ID.Hex()

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &mensaje)
	if err != nil {
		r.Message = "Erroro al decodificar el body " + err.Error()
		return r
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}
	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		r.Message = "Error al insertar registro " + err.Error()
		return r
	}
	if !status {
		r.Message = "No se ha podido insertar el tweet "
		return r
	}

	r.Status = 200
	r.Message = "Tweet insertado con exito."
	return r

}
