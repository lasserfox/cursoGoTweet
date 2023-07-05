package routers

import (
	"context"
	"cursoGoTweet/bd"
	"cursoGoTweet/jwt"
	"cursoGoTweet/models"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"time"
)

func Login(ctx context.Context) models.RespApi {
	var t models.Usuario
	var r models.RespApi
	r.Status = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Usuario y/o contraseña inválidos " + err.Error()
		return r
	}
	if len(t.Email) == 0 {
		r.Message = "El mail no puede ser vacio"
		return r
	}

	userData, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		r.Message = "Usuario y/o contraseña inválidos"
		return r
	}
	jwtKey, err := jwt.GeneroJWT(ctx, userData)
	if err != nil {
		r.Message = "Problemas al generar jwt >" + err.Error()
		return r
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}
	token, err2 := json.Marshal(resp)
	if err2 != nil {
		r.Message = "Problemas al hacer el marshal>" + err2.Error()
		return r
	}
	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(time.Hour * 24),
	}
	cookieString := cookie.String()
	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origins": "*",
			"Set-Cookie":                   cookieString,
		},
	}
	r.Status = 200
	r.Message = string(token)
	r.CustomResp = res
	return r
}
