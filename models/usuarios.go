package models
 
import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)
 
/*Usuario es el modelo de Usuario de la base de MongoDB*/
type Usuario struct {
    /* El ID de MongoDB no es un número, no es un long, ni un string, es un
    objeto de tipo binario de un tipo llamado ObjectID,
    es un slice de bits, que nos llevará a utilizar luego algunas
    funciones especiales por este tipo de datos */
    ID          primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
    Nombre      string              `bson:"nombre" json:"nombre,omitempty"`
    Apellidos   string              `bson:"apellidos" json:"apellidos,omitempty"`
    FechaNacimiento time.Time       `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
    Email       string              `bson:"email" json:"email"`
    Password    string              `bson:"password" json:"password,omitempty"`
    Avatar      string              `bson:"avatar" json:"avatar,omitempty"`
    Banner      string              `bson:"banner" json:"banner,omitempty"`
    Biografia   string              `bson:"biografia" json:"biografia,omitempty"`
    Ubicacion   string              `bson:"ubicacion" json:"ubicacion,omitempty"`
    SitioWeb    string              `bson:"sitioWeb" json:"sitioWeb,omitempty"`
 
}
