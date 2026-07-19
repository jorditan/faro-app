package transaction

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("transaction not found")

type Repository interface {
	Create(transaction Transaction) (Transaction, error)
	FindByID(id int64) (Transaction, error)
	FindAll() ([]Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	Delete(id int64) error
}

type MemoryRepository struct {
	mu     sync.RWMutex
	data   map[int64]Transaction
	nextID int64
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		data:   make(map[int64]Transaction),
		nextID: 1,
	}
}

func (repository *MemoryRepository) Create(transaction Transaction) (Transaction, error) {
	repository.mu.Lock()
	defer repository.mu.Unlock()

	transaction.ID = repository.nextID
	repository.data[transaction.ID] = transaction
	repository.nextID++

	return transaction, nil
}

func (repository *MemoryRepository) FindByID(id int64) (Transaction, error) {
	repository.mu.RLock()
	defer repository.mu.RUnlock()

	transaction, exists := repository.data[id]
	if !exists {
		return Transaction{}, ErrNotFound
	}

	return transaction, nil
}

func (repository *MemoryRepository) FindAll() ([]Transaction, error) {
	repository.mu.RLock()
	defer repository.mu.RUnlock()

	transactions := make([]Transaction, 0, len(repository.data))
	for _, transaction := range repository.data {
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (repository *MemoryRepository) Update(transaction Transaction) (Transaction, error) {
	repository.mu.Lock()
	defer repository.mu.Unlock()

	if _, exists := repository.data[transaction.ID]; !exists {
		return Transaction{}, ErrNotFound
	}

	repository.data[transaction.ID] = transaction

	return transaction, nil
}

func (repository *MemoryRepository) Delete(id int64) error {
	repository.mu.Lock()
	defer repository.mu.Unlock()

	if _, exists := repository.data[id]; !exists {
		return ErrNotFound
	}

	delete(repository.data, id)
	return nil
}package transaction

import "errors"

var ErrNotFound = errors.New("transaction not found")

type Repository interface {
	Create(transaction Transaction) (Transaction, error)
	FindByID(id int64) (Transaction, error)
	FindAll() ([]Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	Delete(id int64) error
}