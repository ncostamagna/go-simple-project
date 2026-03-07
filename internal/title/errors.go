package title

import "errors"

var (
	ErrTitleNotFound = errors.New("title not found")
	ErrNameAndDescriptionRequired = errors.New("name and description are required")
	ErrNoTitlesInDatabase = errors.New("no titles in database")
)

