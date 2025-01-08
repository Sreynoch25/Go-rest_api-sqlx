package auth_service

import (
	"errors"
	"fmt"
	auth_model "marketing/src/models/auth"
	auth_repository "marketing/src/repositeries/auth"
	"marketing/src/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	UserLogin(loginRequest auth_model.UserLogin) (*auth_model.LoginResponse, error)
	generateJWT(user *auth_model.AuthUser) (string, error)
}

type authService struct {
	repo      auth_repository.AuthRepository
	jwtSecret []byte
}

func NewAuthService(repo auth_repository.AuthRepository, jwtSecret string) AuthService {
	return &authService{
		repo:      repo,
		jwtSecret: []byte(jwtSecret),
	}
}

func (s *authService) UserLogin(loginRequest auth_model.UserLogin) (*auth_model.LoginResponse, error) {
	// Get user by email
	auth, err := s.repo.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password: %w", err)
	}

	// Verify password
	if !utils.CheckPasswordHash(loginRequest.Password, auth.Password) {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := s.generateJWT(auth)
	if err != nil {
		return nil, fmt.Errorf("error generating token: %w", err)
	}

	// Update user login information
	err = s.repo.UpdateUserLogin(auth.ID, token)
	if err != nil {
		return nil, fmt.Errorf("error updating user login information: %w", err)
	}

	return &auth_model.LoginResponse{
		Message:    "Login successful",
		Token:      token,
		StatusCode: 200,
		Success:    true,
	}, nil
}

func (s *authService) generateJWT(auth *auth_model.AuthUser) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"user_id":       auth.ID,
		"user_name":     auth.UserName,
		"email":         auth.Email,
		"role_id":       auth.RoleID,
		"login_session": auth.LoginSession,
		"exp":          expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}