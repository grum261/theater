package pgdb

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type pgdb interface {
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	BeginTx(context.Context, pgx.TxOptions) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Begin(context.Context) (pgx.Tx, error)
	Acquire(context.Context) (*pgxpool.Conn, error)
}

type Queries struct {
	db pgdb
}

func newQueries(db pgdb) *Queries {
	return &Queries{
		db: db,
	}
}
