package routers
 
import (
    "encoding/json"
    "net/http"
    "strconv"
 
    "github.com/aguszagame/microblogging/bd"
)
 
/*ListaUsuarios - lee la lista de los usuarios*/
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {
 
    //vamos a capturar los parámetros que vienen en la URL
    //en type indico si quiero los usuarios que yo sigo o los que no sigo
    typeUser := r.URL.Query().Get("type")
    //en page la página de resultados que estoy mostrando
    page := r.URL.Query().Get("page")
    //en search tengo el término de búsqueda que puede o no venir
    search := r.URL.Query().Get("search")
 
    //convierto la pagina que viene como string a numero
    pageTemp, err := strconv.Atoi(page)
    if err != nil {
        http.Error(w, "Debe enviar el parámetro pagina como entero mayor a 0", http.StatusBadRequest)
        return  
    }
 
    pag := int64(pageTemp)
    // llamo a la función de bd con el ID del usuario que esta logueado y lo que vino
    result, status := bd.LeoUsuariosTodos(IDUsuario, pag, search, typeUser)
    if status == false {
        http.Error(w, "Error al leer los usuarios ", http.StatusBadRequest)
        return
    }
 
    //si no encontró registros el status va a venir en true y result en vacío
    //viene todo bien
 
    // establecemos el tipo de Header
    w.Header().Set("Content-type", "application/json")
    // le damos un status created
    w.WriteHeader(http.StatusCreated)
    // le devolvemos la lista de resultados al navegador
    json.NewEncoder(w).Encode(result)
}