package user_model

import "time"

type User struct {
	ID           int        `json:"id" db:"id"`
	LastName     string     `json:"last_name" db:"last_name" validate:"required"`
	FirstName    string     `json:"first_name" db:"first_name" validate:"required"`
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
	StatusID     int        `json:"status_id" db:"status_id"`
	Order        *int       `json:"order" db:"order"`
	CreatedBy    int        `json:"created_by" db:"created_by"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedBy    int        `json:"updated_by" db:"updated_by"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedBy    *int       `json:"deleted_by" db:"deleted_by"`
	DeletedAt    *time.Time `json:"deleted_at" db:"deleted_at"`
}

