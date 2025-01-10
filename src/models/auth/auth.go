package auth_model

type AuthUser struct {
    ID           int    `db:"id" json:"id"`
	UserName     string `json:"user_name" db:"user_name"`
    Email        string `db:"email" json:"email"`
    Password     string `db:"password" json:"password"`
    RoleName     string `db:"role_name" json:"role_name"`
    RoleID       int    `db:"role_id" json:"role_id"`
    IsAdmin      bool   `db:"is_admin" json:"is_admin"`
    LastLogin    *string `db:"last_login" json:"last_login"`
    LoginSession *string `db:"login_session" json:"login_session"`
}

type LoginRequest struct {
    UserName     string `json:"user_name" db:"user_name"`
    Password string   `json:"password" validate:"required"`
}
 

type LoginResponse struct {
    Message    string `json:"message"`
    Token      string `json:"token"`
    StatusCode int    `json:"status_code"`
    Success    bool   `json:"success"`
}