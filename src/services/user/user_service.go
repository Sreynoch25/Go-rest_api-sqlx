package user_service

import (
	user_model "marketing/src/models/user"
	user_repository "marketing/src/repositeries/user"
)

type UserService interface {
	Show() (*user_model.UserReponse, error)
	ShowOne(id int) (*user_model.User, error)
}

type userService struct {
	repo user_repository.UserRepository
}

func NewUserService(repo user_repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Show() (*user_model.UserReponse, error) {
	return s.repo.Show()
}

func (s *userService) ShowOne(id int) (*user_model.User, error) {
	return s.repo.ShowOne(id)
}