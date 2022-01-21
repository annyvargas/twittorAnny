package main

import (
	"log"

	"github.com/annyvargas/twittorAnny/bd"
	"github.com/annyvargas/twittorAnny/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("sin conexion a la bd")
		return
	}
	handlers.Manejadores()
}
