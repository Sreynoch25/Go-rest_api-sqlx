package auth_service

import (
	auth_model "marketing/src/models/auth"
	auth_repository "marketing/src/repositeries/auth"

	"github.com/jmoiron/sqlx"
)

type AuthService interface {
	UserLogin(loginRequest auth_model.LoginRequest) (*auth_model.LoginResponse, error)
}

type authService struct {
	repo auth_repository.AuthRepository
}

func NewAuthService(db *sqlx.DB) AuthService {
	return &authService{
		repo: auth_repository.NewAuthRepository(db),
	}
	
}

func (s *authService) UserLogin(loginRequest auth_model.LoginRequest) (*auth_model.LoginResponse, error) {
	
	return s.repo.Login(loginRequest)
}

