package models
 
/*RespuestaLogin tiene el token que se devuelve con el login*/
type RespuestaLogin struct {
    // cuando colocamos el omitempty es porque en caso de error debe devolver vacío
    Token string `json:"token,omitempty"`
}
