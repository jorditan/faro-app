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
	})
	if err != nil {
		return Response{}, err
	}

	return toResponse(createdCategory), nil
}
