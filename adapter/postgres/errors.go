package postgres

import "errors"

var (
	ErrTitleNotFoundPostgres = errors.New("title not found")
	ErrNoTitlesInPostgres = errors.New("no titles in database")
)

