package data

import (
	"time"

	"github.com/uptrace/bun"
)

type CallModel struct {
	bun.BaseModel `bun:"table:calls,alias:c"`

	ID          int64  `bun:"id,pk"`
	MutationID  int64  `bun:"mutation_id,notnull"`
	SessionID   int64  `bun:"session_id"`
	Body        string `bun:"body"`
	PathParams  string `bun:"path_params"`
	QueryParams string `bun:"query_params"`
	Arguments   string `bun:"arguments"`

	CreatedAt time.Time `bun:"created_at,default:now()"`
}
