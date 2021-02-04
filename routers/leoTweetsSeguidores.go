
package routers
 
import (
    "encoding/json"
    "net/http"
    "strconv"
 
    "github.com/aguszagame/microblogging/bd"
)
 
/*LeoTweetsSeguidores - lee los tweets de todos nuestros seguidores*/
func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {
 
    //Chequeamos que haya venido la página
    if len(r.URL.Query().Get("pagina")) < 1 {
        http.Error(w, "Debe enviar el parámetro pagina", http.StatusBadRequest)
        return
    }
    // Vamos a trabajar con la paginación
    pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
    if err != nil {
        http.Error(w, "Debe enviar el parámetro pagina con un valor mayor a cero", http.StatusBadRequest)
        return
    }
 
    /*Ahora vamos a ejecutar la rutina de base de datos,
    en respuesta nos va a venir todo el documento de json que vamos a
    tener que enviar al HTTP y además un booleano, correcto*/
    respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)
    if correcto == false {
        http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
        return
    }
    //Ya tenemos nuestro documento de respuesta y todo fue ok
    // establecemos el tipo de Header con Json
    w.Header().Set("Content-type", "application/json")
    // le damos un status created
    w.WriteHeader(http.StatusCreated)
    // le devolvemos la respuesta en formato Json que está en respuesta
    json.NewEncoder(w).Encode(respuesta)
}
