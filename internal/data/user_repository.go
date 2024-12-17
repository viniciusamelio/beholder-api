package data

import (
	"beholder/internal/boundaries"
	"beholder/internal/data/models"
	"beholder/internal/mappers"

	"github.com/uptrace/bun"
)

type UserRepository struct {
	crud *CrudRepository
}

func NewUserRepository(db *bun.DB) UserRepository {
	return UserRepository{
		crud: &CrudRepository{
			db: db,
		},
	}
}

func (r *UserRepository) CreateUser(input boundaries.CreateUserInput) (*models.UserModel, *error) {
	user := mappers.CreateUserMapper(input)
	r.crud.Model = user
	err := r.crud.Create(user)

	return user, err
}

func (r *UserRepository) DeleteUser(input boundaries.DeleteUserInput) (*int, error) {
	user := new(models.UserModel)
	user.ID = int64(input.ID)
	r.crud.Model = user

	err := r.crud.Delete(QueryFilter{})

	return &input.ID, *err
}
