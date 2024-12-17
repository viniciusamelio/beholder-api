package models

import (
	"time"

	"github.com/uptrace/bun"
)

type SessionModel struct {
	bun.BaseModel `bun:"table:sessions,alias:s"`

	ID        int64   `bun:"id,pk"`
	ProjectID int64   `bun:"project_id,notnull"`
	UserID    *int64  `bun:"user_id"`
	Slug      *string `bun:"slug"`

	CreatedAt time.Time `bun:"created_at,default:current_timestamp"`
}
