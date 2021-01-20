package handlers

import (
	"log"
	"net/http"
	"os"

    "github.com/gorilla/mux"
    "github.com/aguszagame/microblogging/middlew" 
    "github.com/aguszagame/microblogging/routers"

	"github.com/rs/cors"
)

/*Manejadores: seteo mi puerto, el handler y pongo a escuchar al servidor*/
func Manejadores(){
    router := mux.NewRouter()

    /*Por cada EndPoint vamos a tener un renglon de codigo que permita manejar la funcion correspondiente*/
    router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
    router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")

    PORT := os.Getenv("PORT")
    if PORT == ""{
        PORT = "8080"
    }
    handler := cors.AllowAll().Handler(router)
    log.Fatal(http.ListenAndServe(":"+PORT,handler))
}
