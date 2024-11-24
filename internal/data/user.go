package data

import (
	"time"

	"github.com/uptrace/bun"
)

type UserModel struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID     int64  `bun:"id,pk"`
	Name   string `bun:"name,notnull"`
	Email  string `bun:"email,notnull"`
	Status string `bun:"status,notnull,default:active"`

	CreatedAt time.Time `bun:"created_at,default:now()"`
	UpdatedAt time.Time `bun:"updated_at"`
}
