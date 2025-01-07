package user_repository

import (
	"database/sql"
	"fmt"
	user_model "marketing/src/models/user"
	"strings"

	"github.com/jmoiron/sqlx"
)

/*
 * Author: Noch
 * UserRepository define the interface for user database operations
 */
type UserRepository interface {
	Create(user *user_model.CreateUserRequest) (*user_model.User, error)
	Update(id int, user *user_model.UpdateUserRequest) (*user_model.User, error)
	Show(page, perPage int) (*user_model.UsersResponse, error)
	ShowOne(id int) (*user_model.User, error)
	Delete(id int, deletedBy int) error
}

/*
 * Author: Noch
 * UserRepository implements the UserRepository interface
 */
type userRepository struct {
	db *sqlx.DB // Database connection
}

/*
 * Author: Noch
 * UserRepository create a new repository instance with the provided database connection
 */
func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

/*
 * Author: Noch
 * UserRepository inserts a new user record into the database
  *Params:
 *	-  userReq: Create inserts to be inserted into the database
  *Return:
 *	-  created user object and any  error
*/
func (repo *userRepository) Create(userReq *user_model.CreateUserRequest) (*user_model.User, error) {
	query := `
    INSERT INTO tbl_users (
        user_name, 
        login_id,
        email,
        password,
        role_name,
        role_id,
        is_admin,
        currency_id,
        language_id,
        profile,
        parent_id,
        level,
        status_id,
        "order",
        created_by
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
    ) RETURNING *`

	var user user_model.User
	err := repo.db.Get(&user, query,
		userReq.UserName,
		userReq.LoginID,
		userReq.Email,
		userReq.Password,
		userReq.RoleName,
		userReq.RoleID,
		userReq.IsAdmin,
		userReq.CurrencyID,
		userReq.LanguageID,
		userReq.Profile,
		userReq.ParentID,
		userReq.Level,
		userReq.StatusID,
		userReq.Order,
		userReq.CreatedBy,
	)

	if err != nil {
		// Check specifically for duplicate key violation
		if strings.Contains(err.Error(), "duplicate key value") {
			if strings.Contains(err.Error(), "login_id_key") {
				return nil, fmt.Errorf("login ID already exists")
			}
			if strings.Contains(err.Error(), "email_key") {
				return nil, fmt.Errorf("email already exists")
			}
			return nil, fmt.Errorf("duplicate value not allowed")
		}
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	return &user, nil
}

/*
 * Author: Noch
 *  Update modifies an existing user record
  *Params:
 *	-  id: The ID of the user to be updated
 *	-  userReq: updated user data
  *Return:
 *	-  updated user object and any  error
*/
func (repo *userRepository) Update(id int, userReq *user_model.UpdateUserRequest) (*user_model.User, error) {
	query := `
        UPDATE tbl_users SET
            user_name = $1,
            login_id = $2,
            email = $3,
            role_name = $4,
            role_id = $5,
            is_admin = $6,
            currency_id = $7,
            language_id = $8,
            profile = $9,
            parent_id = $10,
            level = $11,
            status_id = $12,
            "order" = $13,
            updated_by = $14,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = $15 AND deleted_at IS NULL
        RETURNING *`

	var updatedUser user_model.User
	err := repo.db.Get(&updatedUser, query,
		userReq.UserName,
		userReq.LoginID,
		userReq.Email,
		userReq.RoleName,
		userReq.RoleID,
		userReq.IsAdmin,
		userReq.CurrencyID,
		userReq.LanguageID,
		userReq.Profile,
		userReq.ParentID,
		userReq.Level,
		userReq.StatusID,
		userReq.Order,
		userReq.UpdatedBy,
		id,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user with ID %d not found", id)
	}

	if err != nil {

		return nil, fmt.Errorf("error updating user: %v", err)
	}

	return &updatedUser, nil
}

/*
 * Author: Noch
 *   Show retrieves a paginated list of users
  *Params:
 *	-   page: Page number (1-based)
 *	-   perPage: Number of items per page
  *Return:
 *	-  UsersResponse containing users and total count
*/
func (repo *userRepository) Show(page, perPage int) (*user_model.UsersResponse, error) {
	// Calculate offset for pagination
	offset := (page - 1) * perPage

	// Get total count of non-deleted users
	var total int
	countQuery := `SELECT COUNT(*) FROM tbl_users WHERE deleted_at IS NULL`
	err := repo.db.Get(&total, countQuery)
	if err != nil {
		return nil, fmt.Errorf("error counting users: %v", err)
	}

	// Fetch paginated users
	query := `
        SELECT * FROM tbl_users 
        WHERE deleted_at IS NULL 
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2`

	var users []user_model.User
	err = repo.db.Select(&users, query, perPage, offset)
	if err != nil {
		return nil, fmt.Errorf("error fetching users: %v", err)
	}

	return &user_model.UsersResponse{
		Users: users,
		Total: total,
	}, nil
}

/*
 * Author: Noch
 * Show retrieves a single of user by ID
  *Params:
 *	- id: retrieve  user by ID
  *Return:
 *	- User object if found, nil if not found, error if query fails
*/
func (repo *userRepository) ShowOne(id int) (*user_model.User, error) {
	query := `
        SELECT * FROM tbl_users 
        WHERE id = $1 AND deleted_at IS NULL`

	var user user_model.User
	err := repo.db.Get(&user, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %v", err)
	}

	return &user, nil
}

/*
 * Author: Noch
 * Delete performs a soft delete on a user record
  *Params:
 *	- id: ID of user to delete
 *  - deletedBy: ID of user performing the deletion
  *Return:
 *	-error if deletion fails or user not found
*/
func (repo *userRepository) Delete(id int, deletedBy int) error {
	query := `
        UPDATE tbl_users SET
            deleted_by = $1,
            deleted_at = CURRENT_TIMESTAMP
        WHERE id = $2 AND deleted_at IS NULL`

	result, err := repo.db.Exec(query, deletedBy, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking delete result: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("user with ID %d not found", id)
	}

	return nil
}
