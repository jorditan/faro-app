package user

import "time"

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) Create(request CreateRequest) (Response, error) {
	createdUser, err := service.repository.Create(User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return Response{}, err
	}

	return toResponse(createdUser), nil
}

func (service *Service) FindByID(id int64) (Response, error) {
	user, err := service.repository.FindByID(id)
	if err != nil {
		return Response{}, err
	}

	return toResponse(user), nil
}

func (service *Service) FindAll() ([]Response, error) {
	users, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]Response, 0, len(users))
	for _, user := range users {
		responses = append(responses, toResponse(user))
	}

	return responses, nil
}

func (service *Service) Update(id int64, request UpdateRequest) (Response, error) {
	current, err := service.repository.FindByID(id)
	if err != nil {
		return Response{}, err
	}

	if request.Password != "" {
		current.Password = request.Password
	}

	current.Name = request.Name
	current.Email = request.Email

	updatedUser, err := service.repository.Update(current)
	if err != nil {
		return Response{}, err
	}

	return toResponse(updatedUser), nil
}

func (service *Service) Delete(id int64) error {
	return service.repository.Delete(id)
}

func toResponse(user User) Response {
	return Response{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}
