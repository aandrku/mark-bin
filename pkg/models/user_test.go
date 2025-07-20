package models

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

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
			name: "Valid ID: alice",
			id:   1,
			err:  nil,
			// INSERT INTO users (username, email, hashed_password, created) VALUES
			// ('alice',   'alice@example.com',   '$2a$12$C6UzMDM.H6dfI/f/IKcEe.EqDkfQ6VVphX1A0rCwbyXU0JvCk4g2e', '2024-09-15 10:30:00');
			user: User{
				ID:             1,
				Username:       "alice",
				Email:          "alice@example.com",
				HashedPassword: []byte("$2a$12$C6UzMDM.H6dfI/f/IKcEe.EqDkfQ6VVphX1A0rCwbyXU0JvCk4g2e"),
				Created:        time.Date(2024, time.September, 15, 10, 30, 0, 0, time.UTC),
			},
		},
		{
			name: "Valid ID: alice",
			id:   2,
			err:  nil,
			// INSERT INTO users (username, email, hashed_password, created) VALUES
			// ('bob',     'bob@example.com',     '$2a$10$7EqJtq98hPqEX7fNZaFWoO.J8z6NDF8HiK8l2r3L1GxFz0x0miY3C', '2024-10-01 15:45:00');
			user: User{
				ID:             2,
				Username:       "bob",
				Email:          "bob@example.com",
				HashedPassword: []byte("$2a$10$7EqJtq98hPqEX7fNZaFWoO.J8z6NDF8HiK8l2r3L1GxFz0x0miY3C"),
				Created:        time.Date(2024, time.October, 1, 15, 45, 0, 0, time.UTC),
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			db, err := openTestDB(t)
			assert.NilError(t, err)

			m := UserModel{DB: db}

			u, err := m.Get(context.Background(), tc.id)
			assert.Equal(t, tc.user, u)
			assert.Equal(t, tc.err, err)

		})
	}
}
