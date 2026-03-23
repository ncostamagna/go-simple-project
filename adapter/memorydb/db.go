package memorydb

import (
	"sync"
	"github.com/ncostamagna/go-simple-project/domain"
)

type Repository interface {
	Store(title domain.Title) (domain.Title, error)
	GetAll() []domain.Title
	Get(id string) (domain.Title, error)
}

type memoryDB struct {
	database []domain.Title
	dbMu sync.Mutex
}

func NewRepository() Repository {
	return &memoryDB{
		database: make([]domain.Title, 0),
	}
}

func (r *memoryDB) Store(title domain.Title) (domain.Title, error) {
	r.dbMu.Lock()
	defer r.dbMu.Unlock()
	
	r.database = append(r.database, title)
	return title, nil
}

func (r *memoryDB) GetAll() []domain.Title {
	return r.database
}

func (r *memoryDB) Get(id string) (domain.Title, error) {

	if len(r.database) == 0 {
		return domain.Title{}, ErrNoTitlesInMemoryDB
	}

	for _, title := range r.database {
		if title.ID == id {
			return title, nil
		}
	}

	return domain.Title{}, ErrTitleNotFoundMemoryDB
}