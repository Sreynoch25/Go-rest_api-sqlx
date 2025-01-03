package user_repository

import (
	"database/sql"
	"fmt"
	user_model "marketing/src/models/user"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(user *user_model.UserRequest) (*user_model.User, error)
    Update(id int, user *user_model.UserRequest) (*user_model.User, error)
	Show(page, perPage int) (*user_model.UsersResponse, error)
	ShowOne(id int) (*user_model.User, error)
    Delete(id int, deletedBy int) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) Create(userReq *user_model.UserRequest) (*user_model.User, error) {
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
        created_by,
        updated_by
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
    )`

    var user user_model.User

    _, err := repo.db.Exec(query,
        userReq.UserName,
        userReq.LoginID,
        userReq.Email,
        userReq.Password,
        userReq.RoleName,       // role_name from userReq
        userReq.RoleID,         // role_id from userReq
        userReq.IsAdmin,        // is_admin from userReq
        userReq.CurrencyID,     // currency_id from userReq
        userReq.LanguageID,     // language_id from userReq
        userReq.Profile,        // profile from userReq
        userReq.ParentID,       // parent_id from userReq
        userReq.Level,          // level from userReq
        userReq.StatusID,       // status_id from userReq
        userReq.Order,          // order from userReq
        userReq.CreatedBy,      // created_by from userReq
        userReq.UpdatedBy,      // updated_by from userReq
    )

    if err != nil {
        fmt.Printf("error creating user: %v\n", err)
        return nil, fmt.Errorf("error creating user: %v", err)
    }

    return &user, nil
}


// Update performs the update operation in the database
func (repo *userRepository) Update(id int, userReq *user_model.UserRequest) (*user_model.User, error) {
    // Construct update query with all relevant fields
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

    // Execute update query and scan result into User struct
    var updatedUser user_model.User
    err := repo.db.QueryRowx(query,
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
    ).StructScan(&updatedUser)

    if err == sql.ErrNoRows {
        return nil, fmt.Errorf("user with ID %d not found", id)
    }
    if err != nil {
        return nil, fmt.Errorf("error updating user: %v", err)
    }

    return &updatedUser, nil
}

func (repo *userRepository) Show(page, perPage int) (*user_model.UsersResponse, error) {
	offset := (page - 1) * perPage

	var total int
	countQuery := `SELECT COUNT(*) FROM tbl_users WHERE deleted_at IS NULL`
	err := repo.db.Get(&total, countQuery)
	if err != nil {
		return nil, fmt.Errorf("error counting users: %v", err)
	}

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