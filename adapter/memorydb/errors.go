package memorydb

import "errors"

var (
	ErrTitleNotFoundMemoryDB = errors.New("title not found")
	ErrNoTitlesInMemoryDB = errors.New("no titles in database")
)

