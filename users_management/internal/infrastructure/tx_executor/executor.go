package txexecutor

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type txContextMetaType string

type txExecutor struct {
	conn *pgx.Conn
}

const txContextMetaName = "tx"

func NewTxExecutor(conn *pgx.Conn) (*txExecutor, error) {
	return &txExecutor{
		conn: conn,
	}, nil
}

func (e *txExecutor) TxBegin(ctx context.Context, exec func(ctx context.Context) error) error {
	return e.wrapBegin(ctx, exec, pgx.TxOptions{})
}

func (e *txExecutor) TxBeginRepeatableRead(ctx context.Context, exec func(ctx context.Context) error) error {
	return e.wrapBegin(ctx, exec, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
}

func (e *txExecutor) TxBeginSerializable(ctx context.Context, exec func(ctx context.Context) error) error {
	return e.wrapBegin(ctx, exec, pgx.TxOptions{IsoLevel: pgx.Serializable})
}

func (e *txExecutor) wrapBegin(ctx context.Context, exec func(ctx context.Context) error, options pgx.TxOptions) error {
	tx, err := e.conn.BeginTx(ctx, options)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, txContextMetaType(txContextMetaName), tx)
	err = exec(ctx)
	if err != nil {
		tx.Rollback(ctx)
	}

	return tx.Commit(ctx)
}
