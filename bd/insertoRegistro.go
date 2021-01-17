package bd

import (
	"context"
	"time"
	"github.com/aguszagame/microblogging/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

/*InsertoRegistro es la parada final con BD para insertar los datos del usuario*/
func InsertoRegistro (u models.Usuario)(string,bool,error){
	ctx, cancel := context.WithTimeout(context.Background(),15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("usuarios")
	u.Password, _ = EncriptarPassword(u.Password)
	//inserta el usuario
	result, err := col.InsertOne(ctx,u)
	//si da errror
	if err != nil{
		return "", false, err
	}
	//si se inserto bien
	ObjID,_ := result.InsertedID.(primitive.ObjectID)
	//convierto el ObjID a string
	return ObjID.String(), true, nil
}
