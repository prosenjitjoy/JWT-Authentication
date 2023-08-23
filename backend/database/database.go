package database

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Database struct {
	db *pgx.Conn
}

func (d *Database) Close(ctx context.Context) {
	d.db.Close(ctx)
}

func (d *Database) GetDB() *pgx.Conn {
	return d.db
}

func NewDatabase(c context.Context) (*Database, error) {
	ctx, cancel := context.WithTimeout(c, time.Minute)
	defer cancel()

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &Database{db: conn}, nil
}
