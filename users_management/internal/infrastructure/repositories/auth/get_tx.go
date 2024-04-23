package auth_repo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *authRepository) NewTx(ctx context.Context) (pgx.Tx, error) {
	return r.conn.Begin(ctx)
}
