package title

import "errors"

var (
	ErrTitleNotFoundTitle = errors.New("title not found")
	ErrNameAndDescriptionRequired = errors.New("name and description are required")
)

