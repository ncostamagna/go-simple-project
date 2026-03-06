package title

import "errors"

var ErrTitleNotFound = errors.New("title not found")
var ErrNameAndDescriptionRequired = errors.New("name and description are required")

