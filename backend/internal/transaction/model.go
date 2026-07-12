package transaction

import "faro/backend/internal/shared"

type Transaction struct {
	ID         int64
	UserID     int64
	CategoryID *int64
	Amount     float64
	Type       shared.TransactionType
}
