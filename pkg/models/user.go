package models

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID             int
	Username       string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModelInterface interface {
	Insert(username, email, password string) error
	Exists(ID int) (bool, error)
	Authenticate(email, password string) (int, error)
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Insert(ctx context.Context, username, email, password string) error {

	return nil
}

// Exists checks if user with given ID exists in a database.
func (u *UserModel) Exists(ctx context.Context, ID int) (bool, error) {
	var id int

	row := u.DB.QueryRowContext(ctx, "SELECT id FROM users WHERE id=$1", ID)
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (u *UserModel) Authenticate(ctx context.Context, email, password string) (int, error) {

	return 1, nil
}
