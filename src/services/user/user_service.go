// services/user/user_service.go
package user_service

import (
	user_model "marketing/src/models/user"
	user_repository "marketing/src/repositeries/user"

	"github.com/jmoiron/sqlx"
)

type UserService interface {
    Show(page, perPage int) (*user_model.UsersResponse, error)
    ShowOne(id int) (*user_model.User, error)
    Create(userReq *user_model.CreateUserRequest) (*user_model.CreateUserResponse, error)
    Update(id int, userReq *user_model.UpdateUserRequest) (*user_model.UpdateUserResponse, error) 
    Delete(id int, deletedBy int) error
}

type userService struct {
    repo     user_repository.UserRepository
}

func NewUserService(db *sqlx.DB) UserService {
    return &userService{
       repo: user_repository.NewUserRepository(db),
    }
}

func (s *userService) Show(page, perPage int) (*user_model.UsersResponse, error) {
    return s.repo.Show(page, perPage)
}

func (s *userService) ShowOne(id int) (*user_model.User, error) {
    return s.repo.ShowOne(id)
}

func (s *userService) Create(userReq *user_model.CreateUserRequest) (*user_model.CreateUserResponse, error) {
    return s.repo.Create(userReq)
}

func (s *userService) Update(id int, userReq *user_model.UpdateUserRequest) (*user_model.UpdateUserResponse, error) {
    return s.repo.Update(id, userReq)
}

func (s *userService) Delete(id int, deletedBy int) error {
    return s.repo.Delete(id, deletedBy)
}

