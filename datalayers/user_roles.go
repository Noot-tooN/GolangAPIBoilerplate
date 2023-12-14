package datalayers

import (
	"fmt"
	"golangapi/constants"
	"golangapi/models"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRoleDatalayer interface {
	AddRoleForUserByUuid(roleUuid uuid.UUID, userUuid uuid.UUID, gDB *gorm.DB) error
	UpsertRoleForUser(roleName constants.RoleName, userUuid uuid.UUID, gDB *gorm.DB) error

	RemoveRoleForUserByUuid(roleUuid uuid.UUID, userUuid uuid.UUID, gDB *gorm.DB) error
	RemoveRoleForUser(roleName constants.RoleName, userUuid uuid.UUID, gDB *gorm.DB) error

	GetUserRoles(userUuid uuid.UUID, gDB *gorm.DB) ([]models.Role, error)
}

type GormUserRoleDatalayer struct{}

func NewGormUserRoleDatalayer() UserRoleDatalayer {
	return GormUserRoleDatalayer{}
}

func (rdl GormUserRoleDatalayer) GetUserRoles(userUuid uuid.UUID, gDB *gorm.DB) ([]models.Role, error) {
	var roles []models.Role

	res := gDB.
		Table("roles").
		Select("roles.name, roles.uuid").
		Joins("inner join user_roles ur on ur.role_uuid = roles.uuid").
		Where("ur.deleted_at IS NULL").
		Where("ur.user_uuid = ?", userUuid).
		Find(&roles)

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("no rows found")
	}

	return roles, res.Error
}

func (rdl GormUserRoleDatalayer) AddRoleForUserByUuid(roleUuid uuid.UUID, userUuid uuid.UUID, gDB *gorm.DB) error {
	userRole := models.UserRole{
		UserUuid: userUuid,
		RoleUuid: roleUuid,
	}

	res := gDB.Create(&userRole)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("no rows created")
	}

	return nil
}

func (rdl GormUserRoleDatalayer) UpsertRoleForUser(roleName constants.RoleName, userUuid uuid.UUID, gDB *gorm.DB) error {
	RoleUuid := clause.Expr{
		SQL: "(SELECT uuid FROM roles WHERE name = ?)",
		Vars: []interface{}{
			roleName,
		},
	}

	now := time.Now()

	userRole := map[string]interface{}{
		"uuid":       uuid.NewV4(),
		"user_uuid":  userUuid,
		"role_uuid":  RoleUuid,
		"created_at": now,
		"updated_at": now,
	}

	res := gDB.
		Table("user_roles").
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "user_uuid"},
				{Name: "role_uuid"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"deleted_at": nil,
			}),
		}).
		Create(&userRole)

	if res.Error != nil {
		fmt.Println(res.Error)
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("no rows created")
	}

	return nil
}

func (rdl GormUserRoleDatalayer) RemoveRoleForUserByUuid(roleUuid uuid.UUID, userUuid uuid.UUID, gDB *gorm.DB) error {
	res := gDB.Delete(&models.Role{}).Where("user_uuid = ?", userUuid).Where("role_uuid = ?", roleUuid)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("no rows deleted")
	}

	return nil
}

func (rdl GormUserRoleDatalayer) RemoveRoleForUser(roleName constants.RoleName, userUuid uuid.UUID, gDB *gorm.DB) error {
	RoleUuid := clause.Expr{
		SQL: "(SELECT uuid FROM roles WHERE name = ?)",
		Vars: []interface{}{
			roleName,
		},
	}

	res := gDB.
		Table("user_roles").
		Where("role_uuid = ?", RoleUuid).
		Where("user_uuid = ?", userUuid).
		Delete(&models.UserRole{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("no rows deleted")
	}

	return nil
}
