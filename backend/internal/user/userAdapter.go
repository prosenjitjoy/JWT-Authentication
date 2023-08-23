package user

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type adapter struct {
	db *pgx.Conn
}

func NewAdapter(db *pgx.Conn) Adapter {
	return &adapter{db: db}
}

func (a *adapter) CreateUser(ctx context.Context, user *User) (*User, error) {
	query := "INSERT INTO users(username, email, password) values($1, $2, $3) returning id"

	err := a.db.QueryRow(ctx, query, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *adapter) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	rows, err := a.db.Query(ctx, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	if err != nil {
		return nil, err
	}

	return &user, nil
}
