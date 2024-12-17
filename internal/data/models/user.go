package models

import (
	"time"

	"github.com/uptrace/bun"
)

type UserModel struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID             int64  `bun:"id,pk"`
	OrganizationID int64  `bun:"organization_id,notnull"`
	Name           string `bun:"name,notnull"`
	Email          string `bun:"email,notnull"`
	Status         string `bun:"status,notnull,default:active"`

	CreatedAt time.Time  `bun:"created_at,default:current_timestamp"`
	UpdatedAt *time.Time `bun:"updated_at"`
}
