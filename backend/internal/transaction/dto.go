package transaction

import "faro/backend/internal/shared"

type CreateRequest struct {
    UserID     int64                  `json:"userId" binding:"required"`
    CategoryID *int64                 `json:"categoryId"`
    Amount     float64                `json:"amount" binding:"required"`
    Type       shared.TransactionType `json:"type" binding:"required"`
}

type UpdateRequest struct {
    UserID     int64                  `json:"userId" binding:"required"`
    CategoryID *int64                 `json:"categoryId"`
    Amount     float64                `json:"amount" binding:"required"`
    Type       shared.TransactionType `json:"type" binding:"required"`
}

type Response struct {
    ID         int64                  `json:"id"`
    UserID     int64                  `json:"userId"`
    CategoryID *int64                 `json:"categoryId"`
    Amount     float64                `json:"amount"`
    Type       shared.TransactionType `json:"type"`
}