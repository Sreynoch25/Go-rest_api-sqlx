package roles_repository

import (
	"database/sql"
	"fmt"
	role_model "marketing/src/models/role"

	"github.com/jmoiron/sqlx"
)

type RoleRepository interface {
	Show(page, perPage int) (*role_model.RolesResponse, error)
	ShowOne(id int) (*role_model.RoleResponse, error)
	Create(roleReq *role_model.CreateRoleRequest) (*role_model.CreateRoleResponse, error)
	Update(id int, roleReq *role_model.UpdateRoleRequest) (*role_model.RoleResponse, error)
	Delete(id int, deletedBy int) error
}

type roleRepository struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (repo *roleRepository) Show(page, perPage int) (*role_model.RolesResponse, error) {
    // Calculate offset for pagination
    offset := (page - 1) * perPage

    var total int
    countQuery := `SELECT COUNT(*) FROM tbl_roles WHERE deleted_at IS NULL`
    err := repo.db.Get(&total, countQuery)
    if err != nil {
        return nil, fmt.Errorf("error counting roles: %v", err)
    }

    // Fetch paginated roles
    query := `
    SELECT * FROM tbl_roles 
    WHERE deleted_at IS NULL 
    ORDER BY created_at DESC
    LIMIT $1 OFFSET $2`

    roles := []role_model.Role{}
    err = repo.db.Select(&roles, query, perPage, offset)
    if err != nil {
        return nil, fmt.Errorf("error fetching roles: %v", err)
    }

    return &role_model.RolesResponse{
        Roles: roles,
        Total: total,  // Add total to the response
    }, nil
}

func (repo *roleRepository) ShowOne(id int) (*role_model.RoleResponse, error) {

	query := `SELECT * FROM tbl_roles 
	 		  WHERE id = $1 AND deleted_at IS NULL`

	// roleRes := role_model.RoleResponse{}
	roleRes := role_model.RoleResponse{}

	err := repo.db.Get(&roleRes.Role, query, id)
	if err != nil {
		return nil, err
	}
	return &roleRes, nil
}

func (repo *roleRepository) Create(roleReq *role_model.CreateRoleRequest) (*role_model.CreateRoleResponse, error) {
    var exists bool
    checkQuery := `SELECT EXISTS(SELECT 1 FROM tbl_roles WHERE role_name = $1)`
    
    err := repo.db.QueryRow(checkQuery, roleReq.RoleName).Scan(&exists)
    if err != nil {
        return nil, fmt.Errorf("error checking role existence: %v", err)
    }
    
    if exists {
        return nil, fmt.Errorf("role with name '%s' already exists", roleReq.RoleName)
    }

    query := `INSERT INTO tbl_roles (role_name, remark, status_id, "order", created_by)
              VALUES ($1, $2, $3, $4, $5)
              RETURNING id, role_name, remark, status_id, "order", created_by, created_at, updated_by, updated_at, deleted_by, deleted_at`

    var role role_model.Role
    
    err = repo.db.Get(&role, query,
        roleReq.RoleName,
        roleReq.Remark,
        roleReq.StatusID,
        roleReq.Order,
        roleReq.CreatedBy,
    )
    
    if err != nil {
        return nil, fmt.Errorf("error creating role: %v", err)
    }

    return &role_model.CreateRoleResponse{
        Role: role,
    }, nil
}


func (repo *roleRepository) Update(id int, roleReq *role_model.UpdateRoleRequest) (*role_model.RoleResponse, error) {
    query := `
        UPDATE tbl_roles 
        SET role_name = $1, 
            remark = $2, 
            status_id = $3, 
            "order" = $4,
            updated_by = $5,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = $6 AND deleted_at IS NULL
        RETURNING id, role_name, remark, status_id, "order", 
                  created_by, created_at, updated_by, updated_at, 
                  deleted_by, deleted_at`

    var role role_model.Role // Changed to scan into Role directly
    err := repo.db.Get(&role, query,
        roleReq.RoleName,
        roleReq.Remark,
        roleReq.StatusID,
        roleReq.Order,
        roleReq.UpdatedBy,
        id,
    )

    if err == sql.ErrNoRows {
        return nil, fmt.Errorf("role with ID %d not found", id)
    }
    if err != nil {
        return nil, fmt.Errorf("error updating role: %v", err)
    }

    return &role_model.RoleResponse{Role: role}, nil
}

func (repo *roleRepository) Delete(id int, deletedBy int) error {
    query := `
        UPDATE tbl_roles SET
            deleted_by = $1,
            deleted_at = CURRENT_TIMESTAMP
        WHERE id = $2 AND deleted_at IS NULL
    `

    result, err := repo.db.Exec(query, deletedBy, id)
    if err != nil {
        return fmt.Errorf("error deleting role: %v", err)
    }
    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking delete result: %v", err)
    }
    if rows == 0 {
        return fmt.Errorf("role with ID %d not found", id)
    }

    return nil
}

