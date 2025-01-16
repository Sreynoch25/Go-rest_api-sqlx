package role_model

import (
	"time"
)

type Role struct {
	ID        int        `db:"id" json:"id"`
	RoleName  string     `db:"role_name" json:"role_name"`
	Remark    string     `db:"remark" json:"remark"`
	StatusID  int        `db:"status_id" json:"status_id"`
	Order     int        `db:"order" json:"order"`
	CreatedBy *int       `db:"created_by" json:"created_by"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedBy *int       `db:"updated_by" json:"updated_by,omitempty"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedBy *int       `db:"deleted_by" json:"deleted_by,omitempty"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

type CreateRoleRequest struct {
	RoleName  string `db:"role_name" json:"role_name"`
	Remark    string `db:"remark" json:"remark"`
	StatusID  int16  `db:"status_id" json:"status_id"`
	Order     int    `db:"order" json:"order"`
	CreatedBy int    `db:"created_by" json:"created_by"`
}

type CreateRoleResponse struct {
	Role *Role `json:"role"`
}

type UpdateRoleRequest struct {
	RoleName  string `db:"role_name" json:"role_name"`
	Remark    string `db:"remark" json:"remark"`
	StatusID  int16  `db:"status_id" json:"status_id"`
	Order     int    `db:"order" json:"order"`
	UpdatedBy int    `db:"updated_by" json:"updated_by,omitempty"`
}

type RoleResponse struct {
	Role Role `json:"role"`
}

type RolesResponse struct {
	Roles []Role `json:"roles"`
	Total int    `json:"total"`
}



type UpdateRoleResponse struct {
	Role Role `json:"role"`
}
