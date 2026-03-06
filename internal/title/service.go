package title

import (
	"sync"

	"github.com/google/uuid"
	"github.com/ncostamagna/go-simple-project/domain"

)

type Service interface {
	Store(title domain.Title) (domain.Title, error)
	GetAll() []domain.Title
	Get(id string) (domain.Title, error)
}

type service struct {
	database []domain.Title
	dbMu sync.Mutex
}

func New() Service {
	return &service{
		database: make([]domain.Title, 0),
	}
}


func (s *service) Store(title domain.Title) (domain.Title, error) {

	if title.Name == "" || title.Description == "" {
		return domain.Title{}, ErrNameAndDescriptionRequired
	}

	title.ID = uuid.NewString()

	s.dbMu.Lock()
	s.database = append(s.database, title)
	s.dbMu.Unlock()

	return title, nil

}

func (s *service) GetAll() []domain.Title {

	return s.database
}

func (s *service) Get(id string) (domain.Title, error) {


	if len(s.database) == 0 {
		return domain.Title{}, ErrNoTitlesInDatabase
	}

	for _, title := range s.database {
		if title.ID == id {
			return title, nil
		}
	}

	return domain.Title{}, ErrTitleNotFound
}