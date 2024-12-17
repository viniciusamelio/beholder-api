package mappers

import (
	"beholder/internal/boundaries"
	"beholder/internal/data/models"
	"beholder/internal/infra"
)

func CreateUserMapper(input boundaries.CreateUserInput) *models.UserModel {
	user := new(models.UserModel)
	user.ID = int64(infra.GenerateSnowflakeID())
	user.OrganizationID = int64(input.OrganizationID)
	user.Email = input.Email
	user.Name = input.Name
	user.Status = "active"

	return user
}
