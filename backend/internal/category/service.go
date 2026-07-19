package category

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) Create(request CreateRequest) (Response, error) {
	createdCategory, err := service.repository.Create(Category{
		Name: request.Name,
		Type: request.Type,
	})
	if err != nil {
		return Response{}, err
	}

	return toResponse(createdCategory), nil
}

func (service *Service) FindByID(id int64) (Response, error) {
	category, err := service.repository.FindByID(id)
	if err != nil {
		return Response{}, err
	}

	return toResponse(category), nil
}

func (service *Service) FindAll() ([]Response, error) {
	categories, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]Response, 0, len(categories))
	for _, category := range categories {
		responses = append(responses, toResponse(category))
	}

	return responses, nil
}

func (service *Service) Update(id int64, request UpdateRequest) (Response, error) {
	current, err := service.repository.FindByID(id)
	if err != nil {
		return Response{}, err
	}

	current.Name = request.Name
	current.Type = request.Type

	updatedCategory, err := service.repository.Update(current)
	if err != nil {
		return Response{}, err
	}

	return toResponse(updatedCategory), nil
}

func (service *Service) Delete(id int64) error {
	return service.repository.Delete(id)
}

func toResponse(category Category) Response {
	return Response{
		ID:   category.ID,
		Name: category.Name,
		Type: category.Type,
	}
}
