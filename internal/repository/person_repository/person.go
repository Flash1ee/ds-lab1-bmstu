package person_repository

import (
	"errors"

	"gorm.io/gorm"

	"crud-app/internal/domain"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(obj *domain.Person) error {
	return r.db.Create(obj).Error
}

func (r *Repository) GetAll() []domain.Person {
	var temp []domain.Person
	r.db.Find(&temp)
	return temp
}

func (r *Repository) Find(id int) (*domain.Person, error) {
	temp := &domain.Person{}
	err := r.db.Where("id = ?", id).First(temp).Error
	switch {
	case err == nil:
		break
	case errors.Is(err, gorm.ErrRecordNotFound):
		err = domain.RecordNotFound
	default:
		err = domain.UnknownError
	}

	return temp, err
}

func (r *Repository) Delete(id int) error {
	person := &domain.Person{ID: id}
	return r.db.Delete(person).Error
}

func (r *Repository) Patch(obj *domain.Person) (*domain.Person, error) {
	res := r.db.Model(obj).Updates(obj)
	if res.Error != nil {
		return nil, domain.UnknownError
	} else if res.RowsAffected != 1 {
		return nil, domain.RecordNotFound
	}

	return r.Find(obj.ID)
}
