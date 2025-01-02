// services/user/user_service.go
package user_service

import (
    "fmt"
    user_model "marketing/src/models/user"
    user_repository "marketing/src/repositeries/user"
    "github.com/go-playground/validator/v10"
)

type UserService interface {
    Create(userReq *user_model.UserRequest) (*user_model.User, error)
    Update(id int, user *user_model.UserRequest) error
    Delete(id int, deletedBy int) error
    Show(page, perPage int) (*user_model.UsersResponse, error)
    ShowOne(id int) (*user_model.User, error)
}

type userService struct {
    repo     user_repository.UserRepository
    validate *validator.Validate
}

func NewUserService(repo user_repository.UserRepository) UserService {
    return &userService{
        repo:     repo,
        validate: validator.New(),
    }
}

func (s *userService) Create(userReq *user_model.UserRequest) (*user_model.User, error) {
    if err := s.validate.Struct(userReq); err != nil {
        return nil, fmt.Errorf("validation error: %w", err)
    }
    
    return s.repo.Create(userReq)
}

func (s *userService) Update(id int, userReq *user_model.UserRequest) error {
    if err := s.validate.Struct(userReq); err != nil {
        return fmt.Errorf("validation error: %w", err)
    }
    
    return s.repo.Update(id, userReq)
}

func (s *userService) Delete(id int, deletedBy int) error {
    return s.repo.Delete(id, deletedBy)
}

func (s *userService) Show(page, perPage int) (*user_model.UsersResponse, error) {
    return s.repo.Show(page, perPage)
}

func (s *userService) ShowOne(id int) (*user_model.User, error) {
    return s.repo.ShowOne(id)
}