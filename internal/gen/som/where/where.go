// Code generated by github.com/go-surreal/som, DO NOT EDIT.

package where

import (
	"beholder-api/internal/gen/som/internal/lib"
)

func All[M any](filters ...lib.Filter[M]) lib.Filter[M] {
	return lib.All[M](filters)
}

func Any[M any](filters ...lib.Filter[M]) lib.Filter[M] {
	return lib.Any[M](filters)
}
