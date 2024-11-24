package data

import (
	"time"

	"github.com/uptrace/bun"
)

type ResponseModel struct {
	bun.BaseModel `bun:"table:responses,alias:r"`

	ID     int64   `bun:"id,pk"`
	CallID int64   `bun:"call_id,notnull"`
	Data   *string `bun:"body"`
	Status *string `bun:"status"`

	CreatedAt time.Time `bun:"created_at,default:current_timestamp"`
}
