package auth_repository

import (
	"errors"
	"fmt"
	auth_model "marketing/src/models/auth"
	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
    GetUserByEmail(email string) (*auth_model.AuthUser, error)
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

func (repo *authRepository) GetUserByEmail(email string) (*auth_model.AuthUser, error) {
    auth := auth_model.AuthUser{}
    err := repo.db.Get(
        &auth,
        `
        SELECT id, user_name, email, password, role_id, login_session
        FROM tbl_users
        WHERE email = $1
        `, email)
    if err != nil {
        return nil, err
    }
    return &auth, nil
}

// func (repo *authRepository) UpdateUserLogin(userID int, loginSession string) error {
// 	query := `
//     UPDATE tbl_users
//     SET last_login  = :last_login,
//     login_session  = :login_session,
//     updated_at =    CURRENT_TIMESTAMP
//     WHERE id =  :id`

// 	params := map[string]interface{}{
// 		"last_login":    time.Now(),
// 		"login_session": loginSession,
// 		"id":            userID,
// 	}
// 	_, err := repo.db.NamedExec(query, params)

// 	return err
// }

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
