package routers

import (
	"encoding/json"
	"net/http"

	"github.com/annyvargas/twittorAnny/bd"
)

/*VerPerfil permite extraer los valores del perfil*/
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro id", http.StatusBadRequest)
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar buscar el registro"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)

}
