package auth_repo

import (
	"context"

	"users_management/internal/infrastructure/repositories"
	queries "users_management/internal/infrastructure/repositories/auth/queries"

	"github.com/jackc/pgx/v5"
)

type authRepository struct {
	conn    *pgx.Conn
	queries *queries.Queries
}

const txContextMetaName = "tx"

func NewAuthRepository(conn *pgx.Conn) *authRepository {
	return &authRepository{
		conn:    conn,
		queries: queries.New(conn),
	}
}

func (r *authRepository) getQuery(ctx context.Context, opts ...repositories.RepoOption) *queries.Queries {
	repo := &repositories.Repo{}
	for _, opt := range opts {
		opt(repo)
	}

	q := r.queries
	if txTemp := ctx.Value(txContextMetaName); txTemp != nil {
		if tx, ok := txTemp.(pgx.Tx); ok {
			q = q.WithTx(tx)
		}
	}

	return q
}
