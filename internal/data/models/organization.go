package models

import (
	"time"

	"github.com/uptrace/bun"
)

type OrganizationModel struct {
	bun.BaseModel `bun:"table:organizations,alias:o"`

	ID   int64  `bun:"id,pk"`
	Name string `bun:"name,notnull"`

	Members []*UserModel `bun:"rel:has-many,join:id=organization_id"`

	OwnerId int64      `bun:"owner_id"`
	Owner   *UserModel `bun:"rel:belongs-to,join:owner_id=id"`

	CreatedAt time.Time  `bun:"created_at,default:current_timestamp"`
	UpdatedAt *time.Time `bun:"updated_at"`
}
