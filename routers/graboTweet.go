package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/annyvargas/twittorAnny/bd"
	"github.com/annyvargas/twittorAnny/models"
)

/*GraboTweet permite grabar el tweet en la base de datos */
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserId:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar insertar el registro, intente nuevamente "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "no se ha logrado insertar el tweet ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
