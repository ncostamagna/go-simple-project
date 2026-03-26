package title

import "errors"

var (
	// Errores que el Servicio le lanza al mundo (HTTP)
	ErrTitleNotFound = errors.New("title not found")
	ErrNameAndDescriptionRequired = errors.New("name and description are required")
	ErrInvalidDatabaseTarget = errors.New("invalid database target")

	// Errores que el Repositorio le lanza al Servicio (Genérico)
	ErrDBNotFound = errors.New("database: record not found")
)

