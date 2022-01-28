package routers

import (
	"encoding/json"
	"net/http"

	"github.com/annyvargas/twittorAnny/bd"
	"github.com/annyvargas/twittorAnny/models"
)

/*ModificarPerfil modifica el perfil de usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar modificar el registro. reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "no se ha logrado modificar el registro del usuario"+err.Error(), 400)
		return

	}

	w.WriteHeader(http.StatusCreated)

}
