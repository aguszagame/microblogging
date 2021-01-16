package bd
 
import (
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
 
/*MongoCN es el objeto de conexión a la base de datos*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://aguszagame:agus1234@cluster0.iz3nc.mongodb.net/microblogging?retryWrites=true&w=majority")
 
/*ConectarBD es la función que me permite conectar a la base de datos*/
func ConectarBD() *mongo.Client{
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil{
        log.Fatal(err.Error())
        return client
    }
    err = client.Ping(context.TODO(),nil)
    if err != nil{
        log.Fatal(err.Error())
        return client
    }
    log.Println("Conexión exitosa con la BD")
    return client
}
 
/*ChequeoConnection es el ping a la BD*/
func ChequeoConnection() int {
    err := MongoCN.Ping(context.TODO(),nil)
    if err != nil{
        return 0
    }
    return 1
}
