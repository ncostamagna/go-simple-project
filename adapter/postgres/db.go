package postgres

import (
	"gorm.io/gorm"
	"github.com/ncostamagna/go-simple-project/domain"
)

type Repository interface {
	Store(title domain.Title) (domain.Title, error)
	GetAll() []domain.Title
	Get(id string) (domain.Title, error)
}

type postgresDB struct {
	db *gorm.DB
}

func NewRepository(dbGorm *gorm.DB) Repository {
	return &postgresDB{db: dbGorm}
}

func (r *postgresDB) GetAll() []domain.Title {

	var titles []domain.Title
	if err := r.db.Find(&titles).Error; err != nil {
		return []domain.Title{}
	}
	return titles
}
 
func (r *postgresDB) Get(id string) (domain.Title, error) {
	var t domain.Title
    if err := r.db.First(&t, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Title{}, ErrTitleNotFoundPostgres
		}
		return domain.Title{}, err
	}
	return t, nil
}

func (r *postgresDB) Store(title domain.Title) (domain.Title, error) {
	if err := r.db.Create(&title).Error; err != nil {
		return domain.Title{}, err
	}
	return title, nil
}