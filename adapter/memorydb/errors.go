package memorydb

import "errors"

var (
	ErrTitleNotFoundAdapter = errors.New("title not found")
	ErrNoTitlesInDatabase = errors.New("no titles in database")
)

