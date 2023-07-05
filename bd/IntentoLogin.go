package bd

import (
	"cursoGoTweet/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if !encontrado {
		return usu, false
	}

	passBytes := []byte(password)
	passwordDB := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passBytes)

	if err != nil {
		return usu, false
	}
	return usu, true

}
