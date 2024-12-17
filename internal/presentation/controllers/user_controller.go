package controllers

import (
	"beholder/internal/boundaries"
	"beholder/internal/data"
)

type UserController struct {
	UserRepository data.UserRepository
}

func NewUserController(userRepository data.UserRepository) UserController {
	return UserController{
		UserRepository: userRepository,
	}
}

func (c *UserController) CreateUser(input boundaries.CreateUserInput) (*boundaries.CreateUserOutput, error) {
	userModel, err := c.UserRepository.CreateUser(input)
	if err != nil {
		return nil, err
	}

	return &boundaries.CreateUserOutput{
		ID:             int(userModel.ID),
		OrganizationID: int(userModel.OrganizationID),
		Name:           userModel.Name,
		Email:          userModel.Email,
		CreatedAt:      userModel.CreatedAt,
	}, nil
}

func (c *UserController) DeleteUser(input boundaries.DeleteUserInput) (*boundaries.DeleteUserOutput, error) {
	_, err := c.UserRepository.DeleteUser(input)
	if err != nil {
		return nil, err
	}

	return &boundaries.DeleteUserOutput{
		Message: "User Successfully deleted",
	}, nil
}
