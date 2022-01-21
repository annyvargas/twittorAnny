package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la bd*/
var MongoCN = ConectarBD()
var clientsOptions = options.Client().ApplyURI("mongodb+srv://tw:Ik5c98tk7dPkhG3x@cluster0.limaw.mongodb.net/twittor?retryWrites=true&w=majority")

/*ConectarBD es la funcion que me permite conectar la bd*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientsOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("conexion exitosa con la bd")
	return client

}

/*ChequeoConnection es el ping a la bd*/
func ChequeoConnection() int {

	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {

		return 0
	}
	return 1

}
