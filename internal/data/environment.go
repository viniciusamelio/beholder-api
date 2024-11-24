package data

import (
	"time"

	"github.com/uptrace/bun"
)

type EnvironmentModel struct {
	bun.BaseModel `bun:"table:environments,alias:e"`

	ID        int64   `bun:"id,pk"`
	ProjectID int64   `bun:"project_id,notnull"`
	Name      string  `bun:"name,notnull"`
	BaseUrl   string  `bun:"base_url,notnull"`
	Color     *string `bun:"color"`

	CreatedAt time.Time  `bun:"created_at,default:current_timestamp"`
	UpdatedAt *time.Time `bun:"updated_at"`
}
