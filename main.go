package main
 
import (
    "log"
    "github.com/aguszagame/microblogging/handlers"
    "github.com/aguszagame/microblogging/bd"
)
func main(){
    if bd.ChequeoConnection() == 0 {
        log.Fatal("Sin conexión a la BD")
        return
    }
    handlers.Manejadores()
 
}
