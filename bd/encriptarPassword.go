package bd

import "golang.org/x/crypto/bcrypt"

/*encriptarPassword */
func encriptarPassword(pass string) (string, error) {
	//costo eleva al cuadrado y es la cantidad de pasadas que hace sobre el texto para encriptarla, si tiene mas costo mas se demora
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	return string(bytes), err

}
