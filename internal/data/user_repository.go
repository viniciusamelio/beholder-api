package data

import (
	"beholder/internal/boundaries"
	"beholder/internal/infra"
	"context"

	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(input boundaries.CreateUserInput) (*UserModel, error) {
	user := new(UserModel)
	user.ID = int64(infra.GenerateSnowflakeID())
	user.OrganizationID = int64(input.OrganizationID)
	user.Email = input.Email
	user.Name = input.Name
	user.Status = "active"

	_, err := r.db.NewInsert().Model(user).Returning("*").Exec(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
