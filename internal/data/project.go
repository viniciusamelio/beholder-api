package data

import (
	"time"

	"github.com/uptrace/bun"
)

type ProjectModel struct {
	bun.BaseModel `bun:"table:projects,alias:p"`

	ID   int64  `bun:"id,pk"`
	Name string `bun:"name,notnull"`

	OrganizationId int64             `bun:"organization_id"`
	Organization   OrganizationModel `bun:"rel:belongs-to,join:organization_id=id"`

	UserId    int64     `bun:"user_id"`
	CreatedBy UserModel `bun:"rel:belongs-to,join:user_id=id"`

	CreatedAt time.Time `bun:"created_at,default:now()"`
	UpdatedAt time.Time `bun:"updated_at"`
}
