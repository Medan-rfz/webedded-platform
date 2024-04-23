package repositories

import "github.com/jackc/pgx/v5"

type Repo struct {
	Tx pgx.Tx
}

type RepoOption func(*Repo)

func WithTx(tx pgx.Tx) RepoOption {
	return func(r *Repo) {
		r.Tx = tx
	}
}
