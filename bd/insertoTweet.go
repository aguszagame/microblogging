package bd
 
import (
    "context"
    "time"
 
    "github.com/aguszagame/microblogging/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)
 
/*InsertoTweet graba el tweet en la BD*/
func InsertoTweet(t models.GraboTweet) (string, bool, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()
 
    db := MongoCN.Database("microblogging")
    col := db.Collection("tweet")
 
    registro := bson.M{
        "userid":  t.UserID,
        "mensaje": t.Mensaje,
        "fecha":   t.Fecha,
    }
    result, err := col.InsertOne(ctx, registro)
    if err != nil {
        return "", false, err
    }
    ObjID, _ := result.InsertedID.(primitive.ObjectID)
    return ObjID.String(), true, nil
}
