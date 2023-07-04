package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lasserfox/cursoGoTweet/bd"
	"github.com/lasserfox/cursoGoTweet/models"
)

func Registro(ctx context.Context) models.RespApi {
	var t models.Usuario
	var r models.RespApi
	r.Status = 400
	fmt.Println("Entrando a Registro")
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Debe especificar un mail"
		fmt.Println(r.Message)
		return r
	}

	if len(t.Password) < 6 {
		r.Message = "La pass debe tener un minimo de 6 caracteres"
		fmt.Println(r.Message)
		return r
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		r.Message = "Ya existe un usuario registrado con ese email"
		fmt.Println(r.Message)
		return r
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		r.Message = "Ha ocurrdo un error al realizar el insert del usuario " + err.Error()
		return r
	}
	if !status {
		r.Message = "No s ha podido insertar el usuario"
		return r
	}
	r.Status = 200
	r.Message = "Registro OK"
	fmt.Println(r.Message)
	return r
}
