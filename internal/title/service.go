package title

import (	
	"errors"
	"github.com/ncostamagna/go-simple-project/domain"
	"github.com/google/uuid"
	"github.com/ncostamagna/go-simple-project/adapter/postgres"
	"github.com/ncostamagna/go-simple-project/adapter/memorydb"
)


type Service interface {
	Store(t domain.Title) (domain.Title, error)
	GetAll() []domain.Title
	Get(id string) (domain.Title, error)
}

type service struct {
	postgresRepo postgres.Repository
	memoryRepo memorydb.Repository
}

func NewService(p postgres.Repository, m memorydb.Repository) Service {
	return &service{
		postgresRepo: p, 
		memoryRepo: m,
	}
}

func (s *service) Store(t domain.Title) (domain.Title, error) {

	if t.Name == "" || t.Description == "" {
		return domain.Title{}, ErrNameAndDescriptionRequired
	}

	t.ID = uuid.NewString()

	return s.postgresRepo.Store(t)

}

func (s *service) GetAll() []domain.Title {

	return s.postgresRepo.GetAll()
}

func (s *service) Get(id string) (domain.Title, error) {

	t, err := s.postgresRepo.Get(id)
	if err != nil {
		if errors.Is(err, postgres.ErrTitleNotFoundPostgres) {
			return domain.Title{}, ErrTitleNotFound
		}
		return domain.Title{}, err
	}

	return t, nil
}