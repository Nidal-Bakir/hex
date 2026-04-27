package db

import (
	"context"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	adapterInstance *Adapter
)

type Adapter struct {
	ConnPool *pgxpool.Pool
}

func NewAdapter(ctx context.Context, username, password, host, port, database, poolMaxConns string) *Adapter {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	// Reuse Connection
	if adapterInstance != nil {
		return adapterInstance
	}

	conStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable pool_max_conns=%s", username, password, host, port, database, poolMaxConns)
	connectionPool, err := pgxpool.New(ctx, conStr)
	if err != nil {
		log.Fatal(err)
	}

	err = connectionPool.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	adapterInstance = &Adapter{
		ConnPool: connectionPool,
	}

	return adapterInstance
}

func (a Adapter) CloseCon(ctx context.Context) {
	a.ConnPool.Close()
}

func (a Adapter) AppendHistory(ctx context.Context, operation string, answer float64) error {
	query, arg, err := sq.
		Insert("arithmatic_history").
		Columns("operation", "answer").
		Values(operation, answer).
		ToSql()
	if err != nil {
		return err
	}
	_, err = a.ConnPool.Exec(ctx, query, arg...)
	if err != nil {
		return err
	}
	return nil
}
