package routers

import (
	"encoding/json"
	"net/http"
	"github.com/aguszagame/microblogging/bd"
	"github.com/aguszagame/microblogging/models"
)

//Registro es la funcion para crear en la BD el registro de usuario


func Registro (w http.ResponseWriter, r *http.Request){
	var usu models.Usuario
	//Codifica los json que vienen del fronted a formato bson que es el que usa mongoDB
	if err := json.NewDecoder(r.Body).Decode (&usu)
	err != nil {
		//Verifica los datos json si pudieron pasarse a bson
		http.Error(w,"Error en los datos recibidos: " + err.Error(),400)
		return
	}
	/*Si no hubo error con el Body hago las siguientes validcaciones*/
	if len(usu.Email) == 0 {
		http.Error(w,"El email de usuario es requerido",400)
		return
	}
	if len(usu.Password) < 6{
		http.Error(w,"Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}
	_, encontrado,_ := bd.ChequeoYaExisteUsuario(usu.Email)
	if encontrado == true {
		http.Error(w,"Ya existe un usuario registrado con ese email",400)
		return
	}
	_, status, err := bd.InsertoRegistro(usu)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario" + err.Error(),400)
		return
	}
	/*Si llego hasta aca todo anduvo bien*/
	if status == false {
		http.Error(w,"no se ha logrado insertar el registro de usuario", 400 )
		return
	}
	w.WriteHeader(http.StatusCreated)

}