package routers

import (
	"errors"
	"strings"

	"github.com/annyvargas/twittorAnny/bd"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/annyvargas/twittorAnny/models"
)

/*Email valor usado en todos los endpoints*/
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usara en todos los EndPoints*/
var IDUsuario string

/*ProcesoToken proceso token para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	/*if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])*/

	tk = strings.TrimSpace(splitToken[0])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {

			Email = claims.Email
			IDUsuario = claims.ID.Hex()

		}

		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalio")

	}

	return claims, false, string(""), err

}
