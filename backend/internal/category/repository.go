package category

import "errors"

var ErrNotFound = errors.New("category not found")

type Repository interface {
	Create(category Category) (Category, error)
	FindByID(id int64) (Category, error)
	FindAll() ([]Category, error)
	Update(category Category) (Category, error)
	Delete(id int64) error
}
