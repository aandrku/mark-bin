package models

import (
	"context"
	"testing"
	"time"

	"github.com/aandrku/mark-bin/pkg/assert"
)

func TestSnippetModelGet(t *testing.T) {
	tcs := []struct {
		name    string
		id      int
		snippet Snippet
		err     error
	}{
		{
			name:    "Zero ID",
			id:      0,
			snippet: Snippet{},
			err:     ErrNoRecord,
		},
		{
			name:    "Non-existent ID",
			id:      45,
			snippet: Snippet{},
			err:     ErrNoRecord,
		},
		{
			name: "Valid ID",
			id:   1,
			snippet: Snippet{
				ID:      1,
				Title:   "Morning Light",
				Content: `Golden sun rises\nShadows stretch across the field\nDay begins anew`,
				Created: time.Date(2024, time.January, 1, 8, 0, 0, 0, time.UTC),
				Updated: time.Date(2024, time.January, 1, 8, 0, 0, 0, time.UTC),
				UserID:  1,
			},
			err: nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			db, err := openTestDB(t)
			assert.NilError(t, err)

			m := SnippetModel{DB: db}

			s, err := m.Get(context.Background(), tc.id)
			assert.Equal(t, tc.snippet, s)
			assert.Equal(t, tc.err, err)
		})
	}
}

func TestSnippetModelInsert(t *testing.T) {

}

func TestSnippetModelUpdate(t *testing.T) {

}

func TestSnippetModelDelete(t *testing.T) {
	tcs := []struct {
		name string
		id   int
	}{
		{
			name: "Snippet with ID 1",
			id:   1,
		},
		{
			name: "Snippet with ID 2",
			id:   2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			db, err := openTestDB(t)
			assert.NilError(t, err)

			m := SnippetModel{DB: db}

			// user should exist at this point
			_, err = m.Get(context.Background(), tc.id)
			assert.NilError(t, err)

			err = m.Delete(context.Background(), tc.id)
			assert.NilError(t, err)

			// user should not exist after we call delete
			_, err = m.Get(context.Background(), tc.id)
			assert.Equal(t, ErrNoRecord, err)
		})
	}

	t.Run("Delete record that does not exist", func(t *testing.T) {
		db, err := openTestDB(t)
		assert.NilError(t, err)

		id := 45

		m := SnippetModel{DB: db}

		err = m.Delete(context.Background(), id)
		assert.NilError(t, err)
	})

}

func TestSnippetModelLatest(t *testing.T) {

}
