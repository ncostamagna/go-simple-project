package title

import (
	"github.com/ncostamagna/go-simple-project/domain"
	"github.com/google/uuid"
)

type Service interface {
	Store(title domain.Title) (domain.Title, error)
	GetAll() []domain.Title
	Get(id string) (domain.Title, error)
}

type Repository interface {
	Store(title domain.Title) (domain.Title, error)
	GetAll() []domain.Title
	Get(id string) (domain.Title, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}


func (s *service) Store(title domain.Title) (domain.Title, error) {

	if title.Name == "" || title.Description == "" {
		return domain.Title{}, ErrNameAndDescriptionRequired
	}

	title.ID = uuid.NewString()

	return s.repo.Store(title)

}

func (s *service) GetAll() []domain.Title {

	return s.repo.GetAll()
}

func (s *service) Get(id string) (domain.Title, error) {

	return s.repo.Get(id)
}