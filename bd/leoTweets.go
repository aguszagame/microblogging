package bd
 
import (
    "context"
    "log"
    "time"
 
    "github.com/aguszagame/microblogging/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)
 
/*LeoTweets lee los tweets de un perfil, Vamos a devolver un slice, no sabemos la cantidad de tweets que usaremos*/
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()
 
    db := MongoCN.Database("microblogging")
    col := db.Collection("tweet")
 
    // Vamos a crear una variable para los resultados también de tipo del modelo de devuelvoTweets
    // como lo que tengo que devolver
    var resultados []*models.DevuelvoTweets
 
    //En Mongo podemos ver que los tweets tienen un campo userid
    //que lo vamos a comparar con lo que nos viene como parámetro (ID)
    condicion := bson.M{
        "userid": ID,
    }
 
    // Utilizo el paquete option que me permite definir opciones para filtrar y
    // dar un comportamiento a mi consulta de base de datos
 
    // Vamos a trabajar con el modo options en Find
    opciones := options.Find()
 
    // Acá le digo cuantos documentos me traerá a la vez, como máximo
    opciones.SetLimit(20)
 
    //Acá le digo cómo va a ir ordenado lo que traiga, en este caso
    //ordenados por fecha, en orden descendente (se indica con el -1 en Value)
    opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
 
    //Ahora le indico de a cuantos documentos tiene que ir salteando, al principio
    // no saltea, después saltea 20, después 40, y así
    opciones.SetSkip((pagina - 1) * 20)
 
    //Vamos a crear ahora un cursor, que es como un puntero a la tabla
    //donde vamos a ir armando el resultado para devolverlo al router
    //Usamos Find porque queremos que nos encuentre todo lo que cumpla la condicion
    cursor, err := col.Find(ctx, condicion, opciones)
    if err != nil {
        // si no fue satisfactoria la búsqueda en la BD
        log.Fatal(err.Error())
        return resultados, false
    }
 
    // Si todo anduvo bien debo recorrer todos los documentos que encontró
    // el ciclo lo hago con FOR
    // Uso un nuevo contexto vacío para trabajar en él
    for cursor.Next(context.TODO()) {
        // Uso una variable registro con alcance dentro de esta iteración
        // para ir tomando el valor de cada registro leído
        var registro models.DevuelvoTweets
        // leo el registro en cursor
        err := cursor.Decode(&registro)
        if err != nil {
            //si hubo un error devuelvo el resultados vacío
            return resultados, false
        }
        // uso append para agregar un elemento en un slice
        resultados = append(resultados, &registro)
    }
    return resultados, true
}
