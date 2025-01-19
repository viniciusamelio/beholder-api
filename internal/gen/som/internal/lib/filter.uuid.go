// Code generated by github.com/go-surreal/som, DO NOT EDIT.

package lib

import (
	"github.com/google/uuid"

	"beholder-api/internal/gen/som/internal/types"
)

type UUID[M any] struct {
	*Base[M, uuid.UUID, *UUID[M], *Slice[M, uuid.UUID, *UUID[M]]]
}

func NewUUID[M any](key Key[M]) *UUID[M] {
	conv := func(val uuid.UUID) any {
		return types.UUID(val)
	}

	return &UUID[M]{
		Base: NewBaseConv[M, uuid.UUID, *UUID[M], *Slice[M, uuid.UUID, *UUID[M]]](key, conv),
	}
}

type UUIDPtr[M any] struct {
	*UUID[M]
	*Nillable[M]
}

func NewUUIDPtr[M any](key Key[M]) *UUIDPtr[M] {
	return &UUIDPtr[M]{
		UUID:     NewUUID[M](key),
		Nillable: NewNillable[M](key),
	}
}
