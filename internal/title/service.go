package title

import (	
	"errors"
	"github.com/ncostamagna/go-simple-project/domain"
	"github.com/google/uuid"
	"github.com/ncostamagna/go-simple-project/adapter/postgres"
	"github.com/ncostamagna/go-simple-project/adapter/memorydb"
)


type Service interface {
	Store(t domain.Title, dbTarget string) (domain.Title, error)
	GetAll(dbTarget string) ([]domain.Title, error)
	Get(id string, dbTarget string) (domain.Title, error)
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

func (s *service) Store(t domain.Title, dbTarget string) (domain.Title, error) {

	if t.Name == "" || t.Description == "" {
		return domain.Title{}, ErrNameAndDescriptionRequired
	}

	t.ID = uuid.NewString()

	switch dbTarget {
		case "", "postgres":
			return s.postgresRepo.Store(t)
		case "memory":
			return s.memoryRepo.Store(t)
		default:
			return domain.Title{}, ErrInvalidDatabaseTarget
		}
}

func (s *service) GetAll(dbTarget string) ([]domain.Title, error) {
    switch dbTarget {
		case "", "postgres":
			return s.postgresRepo.GetAll(), nil
		case "memory":
			return s.memoryRepo.GetAll(), nil
		default:
			return []domain.Title{}, ErrInvalidDatabaseTarget
		}
}

func (s *service) Get(id string, dbTarget string) (domain.Title, error) {
	switch dbTarget {
		case "", "postgres":
			t, err := s.postgresRepo.Get(id)
			if err != nil {
				if errors.Is(err, postgres.ErrTitleNotFoundPostgres) {
					return domain.Title{}, ErrTitleNotFound
				}
				return domain.Title{}, err
			}
			return t, nil
		case "memory":
			t, err := s.memoryRepo.Get(id)
			if err != nil {
				if errors.Is(err, memorydb.ErrTitleNotFoundMemoryDB) {
					return domain.Title{}, ErrTitleNotFound
				}
				return domain.Title{}, err
			}
			return t, nil
		default:
			return domain.Title{}, ErrInvalidDatabaseTarget
		}
}
