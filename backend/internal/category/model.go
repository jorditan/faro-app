package category

import "faro/backend/internal/shared"

type Category struct {
	ID   int64
	Name string
	Type shared.TransactionType
}
