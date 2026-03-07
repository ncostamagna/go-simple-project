package memorydb

import (
	"sync"

	"github.com/ncostamagna/go-simple-project/domain"
	"github.com/ncostamagna/go-simple-project/internal/title"
)

type repository struct {
	database []domain.Title
	dbMu sync.Mutex
}

func New() title.Repository {
	return &repository{
		database: make([]domain.Title, 0),
	}
}

func (r *repository) Store(title domain.Title) (domain.Title, error) {

	r.dbMu.Lock()
	defer r.dbMu.Unlock()
	
	r.database = append(r.database, title)

	return title, nil
}

func (r *repository) GetAll() []domain.Title {

	result := make([]domain.Title, len(r.database))
	copy(result, r.database)
	
	return result
}

func (r *repository) Get(id string) (domain.Title, error) {


	if len(r.database) == 0 {
		return domain.Title{}, title.ErrNoTitlesInDatabase
	}

	for _, title := range r.database {
		if title.ID == id {
			return title, nil
		}
	}

	return domain.Title{}, title.ErrTitleNotFound
}