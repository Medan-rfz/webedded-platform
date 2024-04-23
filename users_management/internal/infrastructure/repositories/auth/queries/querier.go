// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package auth

import (
	"context"
)

type Querier interface {
	AddRefreshToken(ctx context.Context, arg AddRefreshTokenParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) (int64, error)
	GetAuthByEmail(ctx context.Context, email string) (GetAuthByEmailRow, error)
	GetAuthById(ctx context.Context, id int64) (GetAuthByIdRow, error)
	GetByEmail(ctx context.Context, email string) (GetByEmailRow, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (GetRefreshTokenRow, error)
	InsertCredentials(ctx context.Context, arg InsertCredentialsParams) error
}

var _ Querier = (*Queries)(nil)