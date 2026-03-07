package title

import (
	"errors"
	
	"github.com/ncostamagna/go-simple-project/adapter/memorydb"
	"github.com/ncostamagna/go-simple-project/domain"
	"github.com/google/uuid"
)

type Service interface {
	Store(title domain.Title) (domain.Title, error)
	GetAll() []domain.Title
	Get(id string) (domain.Title, error)
}


type service struct {
	db memorydb.MemoryDB
}

func NewService(db memorydb.MemoryDB) Service {
	return &service{
		db: db,
	}
}


func (s *service) Store(title domain.Title) (domain.Title, error) {

	if title.Name == "" || title.Description == "" {
		return domain.Title{}, ErrNameAndDescriptionRequired
	}

	title.ID = uuid.NewString()

	return s.db.Store(title)

}

func (s *service) GetAll() []domain.Title {

	return s.db.GetAll()
}

func (s *service) Get(id string) (domain.Title, error) {

	title, err := s.db.Get(id)
	if err != nil {
		if errors.Is(err, memorydb.ErrTitleNotFoundAdapter) {
			return domain.Title{}, ErrTitleNotFoundTitle
		}
		return domain.Title{}, err
	}

	return title, nil
}