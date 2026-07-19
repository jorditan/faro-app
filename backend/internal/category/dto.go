package category

import "faro/backend/internal/shared"

type CreateRequest struct {
	Name string                 `json:"name" binding:"required"`
	Type shared.TransactionType `json:"type" binding:"required"`
}

type UpdateRequest struct {
	Name string                 `json:"name" binding:"required"`
	Type shared.TransactionType `json:"type" binding:"required"`
}

type Response struct {
	ID   int64                  `json:"id"`
	Name string                 `json:"name"`
	Type shared.TransactionType `json:"type"`
}
