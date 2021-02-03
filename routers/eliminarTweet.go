package routers
 
import (
    "net/http"
 
    "github.com/aguszagame/microblogging/bd"
)
 
/*EliminarTweet permite borrar un tweet determinado */
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
    ID := r.URL.Query().Get("id")
    if len(ID) < 1 {
        http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
        return
    }
    err := bd.BorroTweet(ID, IDUsuario)
    if err != nil {
        http.Error(w, "Ocurrió un error al intentar borrar el tweet"+err.Error(), http.StatusBadRequest)
        return
    }
 
    /* En la sig línea le decimos al navegador que si va respuesta será del tipo json
       aunque no le mandamos respuesta, es una buena práctica, por si más adelante queremos enviarle una */
    w.Header().Set("Content-type", "application/json")
    // y le damos un status created:
    w.WriteHeader(http.StatusCreated)
}
