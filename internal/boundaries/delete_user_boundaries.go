package boundaries

type DeleteUserInput struct {
	ID int `json:"id"`
}

type DeleteUserOutput struct {
	Message string `json:"message"`
}
