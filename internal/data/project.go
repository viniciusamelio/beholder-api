package data

import (
	"time"

	"github.com/uptrace/bun"
)

type ProjectModel struct {
	bun.BaseModel `bun:"table:projects,alias:p"`

	ID   int64  `bun:"id,pk"`
	Name string `bun:"name,notnull"`

	OrganizationId int64              `bun:"organization_id"`
	Organization   *OrganizationModel `bun:"rel:belongs-to,join:organization_id=id"`

	Environments []*EnvironmentModel `bun:"rel:has-many,join:id=project_id"`
	Mutations    []*MutationModel    `bun:"rel:has-many,join:id=project_id"`

	UserId    int64      `bun:"user_id"`
	CreatedBy *UserModel `bun:"rel:belongs-to,join:user_id=id"`

	CreatedAt time.Time  `bun:"created_at,default:current_timestamp"`
	UpdatedAt *time.Time `bun:"updated_at"`
}
