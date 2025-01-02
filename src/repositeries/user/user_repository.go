package user_repository

import (
    "database/sql"
    "fmt"
    "github.com/jmoiron/sqlx"
    user_model "marketing/src/models/user"
)

type UserRepository interface {
    Create(user *user_model.UserRequest) (*user_model.User, error)
    Update(id int, user *user_model.UserRequest) error
    Delete(id int, deletedBy int) error
    Show(page, perPage int) (*user_model.UsersResponse, error)
    ShowOne(id int) (*user_model.User, error)
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
            user_name, login_id, email, password,
            role_name, role_id, is_admin,
            status_id, created_by, created_at,
            updated_by, updated_at
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9,
            CURRENT_TIMESTAMP, $10, CURRENT_TIMESTAMP
        ) RETURNING *`

    var user user_model.User
	
    err := repo.db.Get(&user, query,  //using Ger() to get a single row
        userReq.UserName,
        userReq.LoginID,
        userReq.Email,
        userReq.Password, 
        userReq.RoleName,
        userReq.RoleID,
        userReq.IsAdmin,
        userReq.StatusID,
        userReq.CreatedBy,
        userReq.UpdatedBy,
    )

    if err != nil {
        return nil, fmt.Errorf("error creating user: %v", err)
    }

    return &user, nil
}

func (repo *userRepository) Update(id int, userReq *user_model.UserRequest) error {
    query := `
        UPDATE tbl_users SET
            user_name = $1,
            login_id = $2,
            email = $3,
            role_name = $4,
            role_id = $5,
            is_admin = $6,
            status_id = $7,
            updated_by = $8,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = $9 AND deleted_at IS NULL`

    result, err := repo.db.Exec(query,
        userReq.UserName,
        userReq.LoginID,
        userReq.Email,
        userReq.RoleName,
        userReq.RoleID,
        userReq.IsAdmin,
        userReq.StatusID,
        userReq.UpdatedBy,
        id,
    )

    if err != nil {
        return fmt.Errorf("error updating user: %v", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking update result: %v", err)
    }

    if rows == 0 {
        return fmt.Errorf("user with ID %d not found", id)
    }

    return nil
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