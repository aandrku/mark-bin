package models

import (
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

func (u *UserModel) Insert(username, email, password string) error {

	return nil
}

func (u *UserModel) Exists(ID int) (bool, error) {

	return true, nil
}

func (u *UserModel) Authenticate(email, password string) (int, error) {

	return 1, nil
}
