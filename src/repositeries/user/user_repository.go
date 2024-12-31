// src/repositories/user/user_repository.go
package user_repository

import (
	"github.com/jmoiron/sqlx"
	user_model "marketing/src/models/user"
)

// UserRepository defines the interface for user-related data access operations.
type UserRepository interface {
	Show() (*user_model.UserReponse, error)
	ShowOne(id int) (*user_model.User, error)
}

// userRepository is a struct that implements the UserRepository interface
type userRepository struct {
	db *sqlx.DB  // Database connection pool
}

// NewUserRepository function will create a new user repository
func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db, // Inject the database connection into the repository
	}
}

/*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             
 *Author:Noch
 * Show retrieves all user records from the database that are not marked as deleted.
*/
func (repo *userRepository) Show() (*user_model.UserReponse, error) {
	users := []user_model.User{} 
	// SQL query to select active users
	query := `
	SELECT id, last_name, first_name, user_name, login_id, email,
		   role_name, role_id, is_admin, login_session, last_login,
		   currency_id, language_id, status_id, "order",
		   created_by, created_at, updated_by, updated_at
	FROM tbl_users 
	WHERE deleted_at IS NULL 
	ORDER BY created_at DESC
`
	err := repo.db.Select(&users, query) // Execute the query and populate the `users` slice
	if err != nil {
		return nil, err // Return nil and the error if the query fails
	}

	response := &user_model.UserReponse{
		User:  users, // Include the retrieved users in the response
		Total: len(users), // Include the total number of users in the response
	}
	return response, nil
}

/*
 *  Author:Noch
 *  ShowOne retrieves a specific user record by its ID
*/
func (repo *userRepository) ShowOne(id int) (*user_model.User, error) {

	user := &user_model.User{}
    // SQL query to select a specific user by ID
	query := `
	SELECT id, last_name, first_name, user_name, login_id, email,
		   role_name, role_id, is_admin, login_session, last_login,
		   currency_id, language_id, status_id, "order",
		   created_by, created_at, updated_by, updated_at
	FROM tbl_users 
	WHERE id = $1 AND deleted_at IS NULL
`
    // Execute the query with the given ID and populate the `user` struct
	err := repo.db.Get(user, query, id)
	if err != nil {   
		return nil, err // Return nil and the error if the query fails
	}

	// Return the retrieved user and no error
	return user, nil
}
