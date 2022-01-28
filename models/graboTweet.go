package models

import "time"

/*GraboTweet es el formato o estructura que tendra nuestro tweet*/
type GraboTweet struct {
	UserId  string    `bson: "userid" json: "userid, omitempty"`
	Mensaje string    `bson: "mensaje" json: "mensaje, omitempty"`
	Fecha   time.Time `bson: "fecha" json: "fecha, omitempty"`
}
