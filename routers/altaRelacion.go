package routers

import (
	"net/http"

	"github.com/annyvargas/twittorAnny/bd"
	"github.com/annyvargas/twittorAnny/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "el parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar insertar relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "no se a logrado insertar relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
