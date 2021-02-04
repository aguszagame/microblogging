package routers
 
import (
    "net/http"
 
    "github.com/aguszagame/microblogging/bd"
    "github.com/aguszagame/microblogging/models"
)
 
/*BajaRelacion - realiza el borrado de la relación entre usuarios*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
    // Obtengo el ID que viene como Parámetro
    ID := r.URL.Query().Get("id")
 
    //Definimos un modelo relación en donde guardaremos lo que vamos a borrar de la bd
    var rel models.Relacion
    //Colocamos como UsuarioID al que tenemos grabado en la variable glogal, que es el logueado
    rel.UsuarioID = IDUsuario
    //Colocamos como UsuarioRelacionID al que viene como parámetro en el query
    rel.UsuarioRelacionID = ID
 
    //Le paso a BorroRelacion el modelo rel que armé
    status, err := bd.BorroRelacion(rel)
    if err != nil {
        http.Error(w, "Ocurrió un error al intentar borrar relacion "+err.Error(), http.StatusBadRequest)
        return
    }
    if status == false {
        http.Error(w, "No se ha logrado borrar la relación. "+err.Error(), http.StatusBadRequest)
        return
    }
    //Si todo estuvo bien con el borrado, mando un StatusCreated
    w.WriteHeader(http.StatusCreated)
}
