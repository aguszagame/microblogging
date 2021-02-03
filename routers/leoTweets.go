package routers
 
import (
    "encoding/json"
    "net/http"
    "strconv"
 
    "github.com/aguszagame/microblogging/bd"
)
 
/*LeoTweets - lee los tweets*/
func LeoTweets(w http.ResponseWriter, r *http.Request) {
 
    ID := r.URL.Query().Get("id")
    if len(ID) < 1 {
        http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
        return
    }
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
    pag := int64(pagina)
    respuesta, correcto := bd.LeoTweets(ID, pag)
    if correcto == false {
        http.Error(w, "error al leer los Tweets", http.StatusBadRequest)
        return
    }
 
    // establecemos el tipo de Header
    w.Header().Set("Content-type", "application/json")
    // le damos un status created
    w.WriteHeader(http.StatusCreated)
    // le devolvemos la respuesta
    json.NewEncoder(w).Encode(respuesta)
}
