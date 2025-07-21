package scheduler

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func InitDB(pool *pgxpool.Pool) {
	db = pool
}

func DB() *pgxpool.Pool {
	return db
}
