package person

import "crud-app/internal/domain"

//go:generate mockgen -source ${GOFILE} -package ${GOPACKAGE}_test -destination mocks_test.go
type storage interface {
	Create(obj *domain.Person) error
	GetAll() []domain.Person
	Find(id int) (*domain.Person, error)
	Delete(id int) error
	Patch(obj *domain.Person) (*domain.Person, error)
}
