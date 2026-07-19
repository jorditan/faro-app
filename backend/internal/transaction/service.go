package transaction

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) Create(request CreateRequest) (Response, error) {
	createdTransaction, err := service.repository.Create(Transaction{
		UserID:     request.UserID,
		CategoryID: request.CategoryID,
		Amount:     request.Amount,
		Type:       request.Type,
	})
	if err != nil {
		return Response{}, err
	}

	return toResponse(createdTransaction), nil
}

func (service *Service) FindByID(id int64) (Response, error) {
	transaction, err := service.repository.FindByID(id)
	if err != nil {
		return Response{}, err
	}

	return toResponse(transaction), nil
}

func (service *Service) FindAll() ([]Response, error) {
	transactions, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]Response, 0, len(transactions))
	for _, transaction := range transactions {
		responses = append(responses, toResponse(transaction))
	}

	return responses, nil
}

func (service *Service) Update(id int64, request UpdateRequest) (Response, error) {
	current, err := service.repository.FindByID(id)
	if err != nil {
		return Response{}, err
	}

	current.UserID = request.UserID
	current.CategoryID = request.CategoryID
	current.Amount = request.Amount
	current.Type = request.Type

	updatedTransaction, err := service.repository.Update(current)
	if err != nil {
		return Response{}, err
	}

	return toResponse(updatedTransaction), nil
}

func (service *Service) Delete(id int64) error {
	return service.repository.Delete(id)
}

func toResponse(transaction Transaction) Response {
	return Response{
		ID:         transaction.ID,
		UserID:     transaction.UserID,
		CategoryID: transaction.CategoryID,
		Amount:     transaction.Amount,
		Type:       transaction.Type,
	}
}