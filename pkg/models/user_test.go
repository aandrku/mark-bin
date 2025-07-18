package models

import (
	"log"
	"os"
	"testing"

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

			exists, err := model.Exists(tc.id)
			if exists != tc.want {
				t.Errorf("got %t, want %t", exists, tc.want)

			}

			if err != nil {
				t.Errorf("err is not nil: %v", err)
			}

			_ = model
		})
	}
}
