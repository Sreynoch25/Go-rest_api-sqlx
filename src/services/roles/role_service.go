package role_service

import (
	role_model "marketing/src/models/role"
	roles_repository "marketing/src/repositeries/roles"
	"github.com/jmoiron/sqlx"
)

type RoleService interface {
	Show(page, perPage int) (*role_model.RolesResponse, error)
	ShowOne(id int) (*role_model.RoleResponse, error)
	Create(roleReq *role_model.CreateRoleRequest) (*role_model.CreateRoleResponse, error)
	Update(id int, roleReq *role_model.UpdateRoleRequest) (*role_model.RoleResponse, error) 
	Delete(id,deleted_by  int) error
}

type roleService struct {
	repo roles_repository.RoleRepository
}

func NewRoleService(db *sqlx.DB) RoleService{
	return &roleService{
		repo: roles_repository.NewRoleRepository(db),
	}
}

func (service *roleService) Show(page, perPage int) (*role_model.RolesResponse, error){
	return service.repo.Show(page, perPage)
}

func (service *roleService) ShowOne(id int) (*role_model.RoleResponse, error){
	return service.repo.ShowOne(id)
}

func (service *roleService) Create(roleReq *role_model.CreateRoleRequest) (*role_model.CreateRoleResponse, error){
	return service.repo.Create(roleReq)
}

func (s *roleService) Update(id int, roleReq *role_model.UpdateRoleRequest) (*role_model.RoleResponse, error) {
    return s.repo.Update(id, roleReq)
}
func (service *roleService) Delete(id,deleted_by  int) error {
	return service.repo.Delete(id, deleted_by)
}