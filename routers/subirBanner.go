package routers
 
import (
    "io"
    "net/http"
    "os"
    "strings"
 
    "github.com/aguszagame/microblogging/bd"
    "github.com/aguszagame/microblogging/models"
)
 
/*SubirBanner sube el Banner al servidor */
func SubirBanner(w http.ResponseWriter, r *http.Request) {
    file, handler, err := r.FormFile("banner")
 
    // Se extrae la extensión del archivo original del banner
    var extension = strings.Split(handler.Filename, ".")[1]
 
    /*En lugar de guardar el nombre que cada usuario le pone al archivo,
      los coloco en una carpeta banners y como nombre le pongo el IDUsuario,
      ya que voy a tener uno por usuario-
      Los archivos se guardan en una carpeta que debe estar previamente creada
      para que todo funcione: carpeta uploads y adentro de ella: carpeta banners*/
    var archivo string = "uploads/banners/" + IDUsuario + "." + extension
 
    // creamos el manejador de archivo con permiso de lectura, escritura y ejecución
    f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        http.Error(w, "Error al subir la imagen "+err.Error(), http.StatusBadRequest)
        return
    }
    /*si no hubo problemas al abrir la imagen, vamos a hacer la copia en f de file,
    además de copiar lo renombra: */
    _, err = io.Copy(f, file)
    if err != nil {
        http.Error(w, "Error al copiar la imagen "+err.Error(), http.StatusBadRequest)
        return
    }
 
    /*Ahora vamos a grabar en la bd el cambio en el campo banner */
    var usuario models.Usuario
    var status bool
    usuario.Banner = IDUsuario + "." + extension
 
    status, err = bd.ModificoRegistro(usuario, IDUsuario)
    if err != nil || status == false {
        http.Error(w, "Error al grabar el banner en la bd "+err.Error(), http.StatusBadRequest)
        return
    }
    w.Header().Set("Content-type", "application/json")
    // y le damos un status created:
    w.WriteHeader(http.StatusCreated)
}
