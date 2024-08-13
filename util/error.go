package util

import "fmt"

var (
	ErrUserNotFound          = fmt.Errorf("usuario no encontrado")
	ErrUserToFollowNotFound  = fmt.Errorf("usuario a seguir no encontrado")
	ErrUsernameAlreadyExists = fmt.Errorf("nombre de usuario ya existente")
	ErrEmailAlreadyExists    = fmt.Errorf("email ya existente")
	ErrInvalidUser           = fmt.Errorf("usuario invalido")
	ErrInvalidPageSize       = fmt.Errorf("tamaño de pagina invalido; debe ser numero mayor a 0")
	ErrInvalidPageNumber     = fmt.Errorf("numero de pagina invalido")
	ErrInvalidDateFormat     = fmt.Errorf("formato de fecha invalido")
	ErrInvalidRequest        = fmt.Errorf("request invalida")
)
