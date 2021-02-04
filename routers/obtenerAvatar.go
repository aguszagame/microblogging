package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/aguszagame/microblogging/bd"
)

/*ObtenerAvatar envía el avatar al http */
func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {

	/*De la URL vamos a obtener nuestro parámetro id del usuario
	  de quien obtenemos el avatar*/
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	/*Busco el perfil por el ID que recibí y lo cargo en un modelo*/
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	/*Intento abrir el archivo*/
	OpenFile, err := os.Open("uploads/avatars/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}
	// Copiamos al ResponseWriter el archivo de la imagen
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la Imagen", http.StatusBadRequest)
	}
	/*En el frontend no van a chequear si hubo un status 201,
	  sino que le haya llegado o no una imagen,
	  por eso no hace falta enviar ningún seteo, ningún status*/
}
