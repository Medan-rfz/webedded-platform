// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package users

import (
	"context"
)

type Querier interface {
	GetById(ctx context.Context, id int64) (User, error)
}

var _ Querier = (*Queries)(nil)