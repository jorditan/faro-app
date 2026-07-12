package category

type CreateRequest struct {
	Name string `json:"name"`
}

type UpdateRequest struct {
	Name string `json:"name"`
}

type Response struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
