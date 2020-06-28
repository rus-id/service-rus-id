package database

import (
	"context"
	"net"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDatabase(ctx context.Context) *pgxpool.Pool {
	c, err := pgxpool.ParseConfig("")
	if err != nil {
		panic("failed to parse postgres config: " + err.Error())
	}

	c.MaxConns = 50
	c.MinConns = 10
	c.ConnConfig.TLSConfig = nil
	c.ConnConfig.PreferSimpleProtocol = true
	c.ConnConfig.RuntimeParams["standard_conforming_strings"] = "on"
	c.ConnConfig.DialFunc = (&net.Dialer{
		KeepAlive: 5 * time.Minute,
		Timeout:   1 * time.Second,
	}).DialContext

	p, err := pgxpool.ConnectConfig(ctx, c)
	if err != nil {
		panic("failed to connect to postgres: " + err.Error())
	}

	return p
}

type contextKey string

const queryer contextKey = "queryer"

type Queryer interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

func NewContext(ctx context.Context, value Queryer) context.Context {
	return context.WithValue(ctx, queryer, value)
}

func FromContext(ctx context.Context) Queryer {
	if queryer, ok := ctx.Value(queryer).(Queryer); ok {
		return queryer
	}

	return nil
}
