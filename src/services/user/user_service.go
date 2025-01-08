// services/user/user_service.go
package user_service

import (
    user_model "marketing/src/models/user"
    user_repository "marketing/src/repositeries/user"
)

type UserService interface {
    Create(userReq *user_model.CreateUserRequest) (*user_model.User, error)
    Update(id int, userReq *user_model.UpdateUserRequest) (*user_model.User, error) 
    Delete(id int, deletedBy int) error
    Show(page, perPage int) (*user_model.UsersResponse, error)
    ShowOne(id int) (*user_model.User, error)
}

type userService struct {
    repo     user_repository.UserRepository
}

func NewUserService(repo user_repository.UserRepository) UserService {
    return &userService{
        repo:     repo,
    }
}

func (s *userService) Create(userReq *user_model.CreateUserRequest) (*user_model.User, error) {
    return s.repo.Create(userReq)
}

func (s *userService) Update(id int, userReq *user_model.UpdateUserRequest) (*user_model.User, error) {
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