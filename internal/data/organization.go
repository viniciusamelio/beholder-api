package data

import (
	"time"

	"github.com/uptrace/bun"
)

type OrganizationModel struct {
	bun.BaseModel `bun:"table:organizations,alias:o"`

	ID   int64  `bun:"id,pk"`
	Name string `bun:"name,notnull"`

	OwnerId int64     `bun:"owner_id"`
	Owner   UserModel `bun:"rel:belongs-to,join:owner_id=id"`

	CreatedAt time.Time `bun:"created_at,default:now()"`
	UpdatedAt time.Time `bun:"updated_at"`
}
