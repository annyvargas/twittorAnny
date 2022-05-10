package routers

import (
	"net/http"

	"github.com/annyvargas/twittorAnny/bd"
	"github.com/annyvargas/twittorAnny/models"
)

// BajaRelacion
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar borrar relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "no se a logrado borrar relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
