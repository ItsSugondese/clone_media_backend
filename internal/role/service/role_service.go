package service

import (
	generic_repo "clone_media/generics/generic-repo"
	"clone_media/internal/role/dto"
	"clone_media/internal/role/model"
	role_navigator "clone_media/internal/role/role-navigator"
	"clone_media/pkg/common/database"
)

func CreateRoleService(dto *dto.RoleRequest) *model.Role {
	tx := database.DB.Begin()

	exists := role_navigator.CheckRoleExistValidationService(dto.Name)

	if exists {
		panic("Role already exists")
	}

	savedRole, saveRoleError := generic_repo.SaveRepo(tx, &model.Role{ID: dto.Name})
	if saveRoleError != nil {
		panic(saveRoleError)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		panic(err)
	}
	return savedRole
}
