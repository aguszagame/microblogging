package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la ruina que me permite encriptar la contraseña recibida*/

func EncriptarPassword(pass string) (string, error){
	costo := 8 //El algoritmo de encriptacion hara (2 elevado al costo) pasadas por el texto 
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass),costo)
	return string(bytes), err
}