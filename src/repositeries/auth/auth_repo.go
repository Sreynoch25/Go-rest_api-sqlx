package auth_repository

import (
	"errors"
	"fmt"
	auth_model "marketing/src/models/auth"
	"marketing/src/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	Login(loginRequest auth_model.LoginRequest) (*auth_model.LoginResponse, error) 
    UpdateUserLogin(userID int, loginSession string) error 
}

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (repo *authRepository) Login(loginRequest auth_model.LoginRequest) (*auth_model.LoginResponse, error) {
	auth := auth_model.AuthUser{}
	err := repo.db.Get(
		&auth,
		`SELECT id, user_name, password
		FROM tbl_users
		WHERE user_name = $1`,
		loginRequest.UserName,
	)
	if err != nil {
		return nil, fmt.Errorf("invalid Username or password: %w", err)
	}

	// Verify password
	if !utils.CheckPasswordHash(loginRequest.Password, auth.Password) {
		return nil, errors.New("invalid Username or password")
	}

	//Generate UUID for login session
	sessionUUID := uuid.New().String()

	// Generate JWT token
	token, err := repo.generateJWT(&auth, sessionUUID)
	if err != nil {
		return nil, fmt.Errorf("error generating token: %w", err)
	}

	// Update user's login information
	err = repo.UpdateUserLogin(auth.ID, sessionUUID)
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


func (repo *authRepository) UpdateUserLogin(userID int, loginSession string) error {
	query := `
        UPDATE tbl_users
        SET 
            last_login = CURRENT_TIMESTAMP,
            login_session = :login_session,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = :id
    `

	params := map[string]interface{}{
		"login_session": loginSession,
		"id":           userID,
	}

	result, err := repo.db.NamedExec(query, params)
	if err != nil {
		return fmt.Errorf("failed to execute update: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return errors.New("no user found with given ID")
	}

	return nil

}

func (repo *authRepository) generateJWT(auth *auth_model.AuthUser, loginSession string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"user_id":       auth.ID,
		"user_name":     auth.UserName,
		"email":         auth.Email,
		"role_id":       auth.RoleID,
		"login_session": loginSession,
		"exp":          expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}