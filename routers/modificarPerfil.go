package routers
 
import (
    "encoding/json"
    "net/http"
 
    "github.com/aguszagame/microblogging/bd"
    "github.com/aguszagame/microblogging/models"
)
 
/*ModificarPerfil - modifica el perfil de usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
 
    var usu models.Usuario
 
    err := json.NewDecoder(r.Body).Decode(&usu)
    if err != nil {
        // es un Json mal construído
        http.Error(w, "Datos incorrectos "+err.Error(), 400)
        return
    }
 
    var status bool
 
    status, err = bd.ModificoRegistro(usu, IDUsuario)
    //IDUsuario es la variable global que seteamos antes con el ID
    if err != nil {
        http.Error(w, "Ocurrió un error al intentar modificar el registro. Intente nuevamente "+err.Error(), 400)
        return
    }
    if status == false {
        http.Error(w, "No se ha logrado modificar el registro del usuario "+err.Error(), 400)
        return
    }
 
    w.WriteHeader(http.StatusCreated)
}
