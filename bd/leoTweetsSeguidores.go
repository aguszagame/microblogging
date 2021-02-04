package bd
 
import (
    "context"
    "time"
 
    "github.com/aguszagame/microblogging/models"
    "go.mongodb.org/mongo-driver/bson"
)
 
/*LeoTweetsSeguidores lee los tweets de mis seguidores*/
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()
 
    db := MongoCN.Database("microblogging")
    col := db.Collection("relacion")
    /*Para esta rutina nuestra tabla principal es la relación, de ahí sacamos la
      información básica de con quién estoy relacionado. Luego vamos a
      unir esta tabla relacion conlos tweets*/
 
    //Para hacer el skip de registros
    skip := (pagina - 1) * 20
 
    /*Vamos a crear un slice condiciones de tipo bson.M que por ahora
      tenga 0 elementos, para no crearlo con tamaño de más*/
    condiciones := make([]bson.M, 0)
 
    /*A condiciones le voy a ir agregando lo necesario que son bson:M, teniendo
    en cuenta la sintaxis del framework Aggregate. Lo primero es un
    comando $match para que busque el usuario que tiene el ID que me llega como parámetro
    en mi relación*/
    condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
    /*Ahora con el comando $lookup voy a unir la tabla relacion con la de tweets,
    tiene 4 parámetros necesarios para unir tablas de MongoDB:
    "from": con la tabla queremos unir la tabla relacion,
    "localField": por qué campo local queremos unirlas,
    "foreignField": con que campo externo quiero unirla,
    "as": un alias de como queremos llamar esa tabla, la llamamos igual
 
    */
    condiciones = append(condiciones, bson.M{
        "$lookup": bson.M{
            "from":         "tweet",
            "localField":   "usuariorelacionid",
            "foreignField": "userid",
            "as":           "tweet",
        }})
 
    /*El comando $unwind hace que no me venga todo como maestro-detalle,
    con ese comando todos los documentos vienen iguales y es más fácil procesar
    la información, aunque venga info repetida*/
    condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
 
    /*Ahora con el comando $sort le indicamos el orden, con la condición
    que le decimos que sea por el campo fecha, en orden descendente, por eso -1*/
    condiciones = append(condiciones, bson.M{"$sort": bson.M{"fecha": -1}})
 
    /*con el comando $skip le decimos de a cuanto hace el salto*/
    condiciones = append(condiciones, bson.M{"$skip": skip})
 
    /*Con el comando $limit le decimos cuantos va a leer */
    condiciones = append(condiciones, bson.M{"$limit": 20})
 
    /*Ahora creamos el cursor con la función Aggregate del framework nuevo,
    que se ejecuta en la bd MongoDB y me trae el cursor.
    Ese cursor directamente ya tiene todo lo que necesitamos para procesar.*/
    cursor, err := col.Aggregate(ctx, condiciones)
 
    /*Creamos un slice para cargar los resultados*/
    var result []models.DevuelvoTweetsSeguidores
 
    /*Con el método All de cursor se procesa todo el cursor, todo junto
    y lo devuelve en result*/
    err = cursor.All(ctx, &result)
    if err != nil {
        return result, false
    }
    return result, true
}
