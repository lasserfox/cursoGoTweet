package jwt

import (
	models2 "cursoGoTweet/models"
	"errors"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/lasserfox/cursoGoTweet/models"
	"strings"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string, JWTSign string) (*models2.Claim, bool, string, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string("Error al procesar el token"), errors.New("formato de token inválido")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		// Rutina que chequea con la BD
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token Inválido")
	}
	return &claims, false, string(""), nil

}
