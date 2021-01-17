package bd

import (
	"context"
	"time"
	"github.com/aguszagame/microblogging/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya esta en la bd*/
func ChequeoYaExisteUsuario(email string)(models.Usuario,bool,string){
	
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	db := MongoCN.Database("microblogging")
	col := db.Collection("usuarios")

	//M es una funcion del paquete bson que formateo o mapea a bson lo que recibe como json
	condicion := bson.M{"email": email}

	//en la variable resultado voy a modelar un usuario
	var resultado models.Usuario

	//FindOne me devuelve un solo registro que cumple con la condicion
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado,false,ID
	}
	return resultado, true,ID

}