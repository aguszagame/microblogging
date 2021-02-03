package bd
 
import (
    "context"
    "time"
 
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)
 
/*BorroTweet borra un tweet determinado*/
func BorroTweet(ID string, UserID string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()
 
    db := MongoCN.Database("microblogging")
    col := db.Collection("tweet")
 
    // convierto el string ID que viene como parámetro en hexadecimal a un ObjID
    objID, _ := primitive.ObjectIDFromHex(ID)
 
    condicion := bson.M{
        "_id":    objID,
        "userid": UserID,
    }
    _, err := col.DeleteOne(ctx, condicion)
    return err
    // directamente retorno err que puede ser un error que dio el borrado o nil si todo estuvo ok
}
