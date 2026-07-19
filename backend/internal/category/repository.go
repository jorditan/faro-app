package category

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("category not found")

type Repository interface {
	Create(category Category) (Category, error)
	FindByID(id int64) (Category, error)
	FindAll() ([]Category, error)
	Update(category Category) (Category, error)
	Delete(id int64) error
}

type MemoryRepository struct {
	mu     sync.RWMutex
	data   map[int64]Category
	nextID int64
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		data:   make(map[int64]Category),
		nextID: 1,
	}
}

func (repository *MemoryRepository) Create(category Category) (Category, error) {
	repository.mu.Lock()
	defer repository.mu.Unlock()

	category.ID = repository.nextID
	repository.data[category.ID] = category
	repository.nextID++

	return category, nil
}

func (repository *MemoryRepository) FindByID(id int64) (Category, error) {
	repository.mu.RLock()
	defer repository.mu.RUnlock()

	category, exists := repository.data[id]
	if !exists {
		return Category{}, ErrNotFound
	}

	return category, nil
}

func (repository *MemoryRepository) FindAll() ([]Category, error) {
	repository.mu.RLock()
	defer repository.mu.RUnlock()

	categories := make([]Category, 0, len(repository.data))
	for _, category := range repository.data {
		categories = append(categories, category)
	}

	return categories, nil
}

func (repository *MemoryRepository) Update(category Category) (Category, error) {
	repository.mu.Lock()
	defer repository.mu.Unlock()

	if _, exists := repository.data[category.ID]; !exists {
		return Category{}, ErrNotFound
	}

	repository.data[category.ID] = category

	return category, nil
}

func (repository *MemoryRepository) Delete(id int64) error {
	repository.mu.Lock()
	defer repository.mu.Unlock()

	if _, exists := repository.data[id]; !exists {
		return ErrNotFound
	}

	delete(repository.data, id)
	return nil
}package category

import "errors"

var ErrNotFound = errors.New("category not found")

type Repository interface {
	Create(category Category) (Category, error)
	FindByID(id int64) (Category, error)
	FindAll() ([]Category, error)
	Update(category Category) (Category, error)
	Delete(id int64) error
}
