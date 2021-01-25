package routers
 
import (
    "errors"
    "strings"
 
    jwt "github.com/dgrijalva/jwt-go"
    "github.com/aguszagame/microblogging/bd"
    "github.com/aguszagame/microblogging/models"
)
 
/* Voy a crear dos variables globales a ProcesoToken para que pueda utilizarlas dentro y fuera
porque serán públicas, aprovecho para setear ahí los valores, una sola vez y luego los utilizo en todo el proceso*/
/*Email - valor del email usado en todos los endpoints*/
var Email string
 
/*IDUsuario es el ID devuelto del modelo, que se usará en todos los endpoints*/
var IDUsuario string
 
/*ProcesoToken - proceso token para extraer sus valores
Es de las funciones más importantes por la cantidad de veces que la vamos a ejecutar,
porque valida el token y dice si la credencial y los privilegios son válidos.
En Go si una función tiene varios parámetros y entre ellos un error, hay que ponerlo al final*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error){
    miClave := []byte("SkillFactoryGo_Avalith")
    // creo una variable claims de tipo claim del models, se indica como puntero
    // porque la estructura en donde vuelca el token debe ser un puntero
    claims := &models.Claim{}
    // El token que viene en el Header comienza con la palabra Bearer, es un estándar, no es parte del token en sí
    splitToken := strings.Split(tk, "Bearer") // aquí separo la palabra Bearer del resto
    if len(splitToken) != 2 {
        // tiene que devolver dos elementos
        return claims, false, string(""), errors.New("formato de token invalido")
        // el error se crea con un mensaje que no puede tener ni mayúsculas, ni tildes, ni símbolos
    }
    tk = strings.TrimSpace(splitToken[1])
    tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token)(interface{}, error){
        // el tercer parámetro es una función anónima que recibe un token y resuelve todo ahí validando el token
        return miClave, nil
    })
    if err == nil { // el token fue válido, hay que ver si el mail que viene en el token es válido
        _,encontrado,_ := bd.ChequeoYaExisteUsuario(claims.Email)
        if encontrado == true{
            Email = claims.Email
            IDUsuario = claims.ID.Hex()
        }
        return claims, encontrado, IDUsuario, nil
    }
    if !tkn.Valid {
        return claims, false, string(""), errors.New("token invalido")
    }
	return claims, false, string(""), err
	}