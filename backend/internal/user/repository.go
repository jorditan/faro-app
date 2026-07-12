package user

import (
	"errors"
	"sync"
	"time"
)

var ErrNotFound = errors.New("user not found")

type Repository interface {
	Create(user User) (User, error)
	FindByID(id int64) (User, error)
	FindAll() ([]User, error)
	Update(user User) (User, error)
	Delete(id int64) error
}

type MemoryRepository struct {
	mu     sync.RWMutex
	data   map[int64]User
	nextID int64
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		data:   make(map[int64]User),
		nextID: 1,
	}
}

func (repository *MemoryRepository) Create(user User) (User, error) {
	repository.mu.Lock()
	defer repository.mu.Unlock()

	now := time.Now()
	user.ID = repository.nextID
	user.CreatedAt = now
	user.UpdatedAt = now
	repository.data[user.ID] = user
	repository.nextID++

	return user, nil
}

func (repository *MemoryRepository) FindByID(id int64) (User, error) {
	repository.mu.RLock()
	defer repository.mu.RUnlock()

	user, exists := repository.data[id]
	if !exists {
		return User{}, ErrNotFound
	}

	return user, nil
}

func (repository *MemoryRepository) FindAll() ([]User, error) {
	repository.mu.RLock()
	defer repository.mu.RUnlock()

	users := make([]User, 0, len(repository.data))
	for _, user := range repository.data {
		users = append(users, user)
	}

	return users, nil
}

func (repository *MemoryRepository) Update(user User) (User, error) {
	repository.mu.Lock()
	defer repository.mu.Unlock()

	current, exists := repository.data[user.ID]
	if !exists {
		return User{}, ErrNotFound
	}

	user.CreatedAt = current.CreatedAt
	user.UpdatedAt = time.Now()
	repository.data[user.ID] = user

	return user, nil
}

func (repository *MemoryRepository) Delete(id int64) error {
	repository.mu.Lock()
	defer repository.mu.Unlock()

	if _, exists := repository.data[id]; !exists {
		return ErrNotFound
	}

	delete(repository.data, id)
	return nil
}
