package models

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/aandrku/mark-bin/pkg/assert"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../.env.test"); err != nil {
		log.Printf("failed to load .env.test: %v", err)
	}

	os.Exit(m.Run())
}

func TestUserModelExists(t *testing.T) {

	tcs := []struct {
		name string
		id   int
		want bool
	}{
		{
			name: "Valid ID",
			id:   1,
			want: true,
		},
		{
			name: "Zero ID",
			id:   0,
			want: false,
		},
		{
			name: "Non-existent ID",
			id:   10,
			want: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			db, err := openTestDB(t)
			if err != nil {
				t.Fatalf("failed to connect to test db: %v", err)
			}

			model := UserModel{DB: db}

			got, err := model.Exists(context.Background(), tc.id)

			assert.Equal(t, tc.want, got)
			assert.NilError(t, err)

			_ = model
		})
	}
}

func TestUserModelGet(t *testing.T) {
	tcs := []struct {
		name string
		id   int
		err  error
		user User
	}{
		{
			name: "Non-existent ID",
			id:   4612,
			err:  ErrNoRecord,
			user: User{},
		},
		{
			name: "Zero ID",
			id:   0,
			err:  ErrNoRecord,
			user: User{},
		},
		{
			name: "Valid ID",
			id:   1,
			err:  nil,
			user: User{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			db, err := openTestDB(t)
			assert.NilError(t, err)

			m := UserModel{DB: db}

			u, err := m.Get(tc.id)
			assert.Equal(tc.user, u)
			assert.Equal(tc.err, err)

		})
	}
}
