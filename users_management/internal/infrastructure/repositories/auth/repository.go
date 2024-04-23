package auth_repo

import (
	"users_management/internal/infrastructure/repositories"
	queries "users_management/internal/infrastructure/repositories/auth/queries"

	"github.com/jackc/pgx/v5"
)

type authRepository struct {
	conn    *pgx.Conn
	queries *queries.Queries
}

func NewAuthRepository(conn *pgx.Conn) *authRepository {
	return &authRepository{
		conn:    conn,
		queries: queries.New(conn),
	}
}

func (r *authRepository) getQuery(opts ...repositories.RepoOption) *queries.Queries {
	repo := &repositories.Repo{}
	for _, opt := range opts {
		opt(repo)
	}

	q := r.queries
	if repo.Tx != nil {
		q = q.WithTx(repo.Tx)
	}

	return q
}
