// models/user/user_model.go
package user_model

import "time"

//struct for user
type User struct {
    ID           int        `json:"id" db:"id"`
    UserName     string     `json:"user_name" db:"user_name" validate:"required"`
    LoginID      string     `json:"login_id" db:"login_id" validate:"required"`
    Email        string     `json:"email" db:"email" validate:"required,email"`
    Password     string     `json:"password,omitempty" db:"password" validate:"required,min=6"`
    RoleName     string     `json:"role_name" db:"role_name" validate:"required"`
    RoleID       int        `json:"role_id" db:"role_id"`
    IsAdmin      bool       `json:"is_admin" db:"is_admin"`
    LoginSession *string    `json:"login_session" db:"login_session"`
    LastLogin    *time.Time `json:"last_login" db:"last_login"`
    CurrencyID   *int       `json:"currency_id" db:"currency_id"`
    LanguageID   *int       `json:"language_id" db:"language_id"`
    Profile      *string    `json:"profile" db:"profile"`
    ParentID     *int       `json:"parent_id" db:"parent_id"`
    Level        string     `json:"level" db:"level"`
    StatusID     int        `json:"status_id" db:"status_id"`
    Order        *int       `json:"order" db:"order"`
    CreatedBy    int        `json:"created_by" db:"created_by"`
    CreatedAt    time.Time  `json:"created_at" db:"created_at"`
    UpdatedBy    int        `json:"updated_by" db:"updated_by"`
    UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
    DeletedBy    *int       `json:"deleted_by" db:"deleted_by"`
    DeletedAt    *time.Time `json:"deleted_at" db:"deleted_at"`
}

//struct for user request
type UserRequest struct {
    UserName    string  `json:"user_name" validate:"required"`
    LoginID     string  `json:"login_id" validate:"required"`
    Email       string  `json:"email" validate:"required,email"`
    Password    string  `json:"password" validate:"required,min=6"`
    RoleName    string  `json:"role_name" validate:"required"`
    RoleID      int     `json:"role_id" validate:"required"`
    IsAdmin     bool    `json:"is_admin"`
    CurrencyID  *int    `json:"currency_id"`
    LanguageID  *int    `json:"language_id"`
    Profile     *string `json:"profile"`
    ParentID    *int    `json:"parent_id"`
    Level       string  `json:"level"`
    StatusID    int     `json:"status_id" validate:"required"`
    CreatedBy   int     `json:"created_by" validate:"required"`
    UpdatedBy   int     `json:"updated_by" validate:"required"`
}

//struct for user response
type UserResponse struct {
    User *User `json:"user"`
}

//struct for users response
type UsersResponse struct {
    Users []User `json:"users"`
    Total int    `json:"total"`
}