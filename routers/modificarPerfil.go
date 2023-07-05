package routers

import (
	"context"
	"cursoGoTweet/bd"
	"cursoGoTweet/models"
	"encoding/json"
)

func ModificarPerfil(ctx context.Context, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	var t models.Usuario
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Datos Incorrectos " + err.Error()
	}
	var status bool
	status, err = bd.ModificoRegistro(t, claim.ID.Hex())
	if err != nil {
		r.Message = "Error al intentar modificar el registro. " + err.Error()
		return r
	}
	if !status {
		r.Message = "No se ha podido modificar el registro del usuario. " + err.Error()
		return r
	}
	r.Status = 200
	r.Message = "Modificación de perfil OK."
	return r
}
