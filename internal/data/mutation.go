package data

import (
	"time"

	"github.com/uptrace/bun"
)

type MutationModel struct {
	bun.BaseModel `bun:"table:mutations,alias:m"`

	ID          int64   `bun:"id,pk"`
	ProjectID   int64   `bun:"project_id,notnull"`
	Name        string  `bun:"name,notnull"`
	Slug        *string `bun:"slug"`
	Description *string `bun:"description"`
	Path        *string `bun:"path"`

	Calls []*CallModel `bun:"rel:has-many,join:id=mutation_id"`

	CreatedAt time.Time  `bun:"created_at,default:current_timestamp"`
	UpdatedAt *time.Time `bun:"updated_at"`
}
