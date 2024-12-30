// service/user_service.go
package user_service

import (
	user_model "marketing/src/models/user"
	user_repository "marketing/src/repositeries/user"
)

type UserService interface {
    GetAll() ([]user_model.User, error)
    GetByID(id int) (*user_model.User, error)
}

type userService struct {
    repo user_repository.UserRepository
}
																	
func NewUserService(repo user_repository.UserRepository) UserService {
    return &userService{
        repo: repo, // Inject the user repository into the service
    }
}

func (s *userService) GetAll() ([]user_model.User, error) {
    return s.repo.GetAll()
}

func (s *userService) GetByID(id int) (*user_model.User, error) {
    return s.repo.GetByID(id)
}
