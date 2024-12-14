package boundaries

import "time"

type CreateUserInput struct {
	OrganizationID int    `json:"organization_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
}

type CreateUserOutput struct {
	ID             int       `json:"id"`
	OrganizationID int       `json:"organization_id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"created_at"`
}
