package routers
 
import (
    "net/http"
 
    "github.com/aguszagame/microblogging/bd"
    "github.com/aguszagame/microblogging/models"
)
 
/*AltaRelacion - realiza el registro de la relación entre usuarios*/
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
    // Obtengo el ID que viene como Parámetro
    ID := r.URL.Query().Get("id")
    if len(ID) < 1 {
        http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
        return
    }
    //Definimos un modelo relación en donde guardaremos lo que vamos a grabar en la bd
    var rel models.Relacion
    //Colocamos como UsuarioID al que tenemos grabado en la variable glogal, que es el logueado
    rel.UsuarioID = IDUsuario
    //Colocamos como UsuarioRelacionID al que viene como parámetro en el query
    rel.UsuarioRelacionID = ID
 
    status, err := bd.InsertoRelacion(rel)
    //si hay un error
    if err != nil {
        http.Error(w, "Ocurrió un error al intentar insertar relación, intentelo nuevamente. "+err.Error(), http.StatusBadRequest)
        return
    }
    if status == false {
        http.Error(w, "No se ha logrado insertar la relación. "+err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
}
