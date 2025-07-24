package models

import (
	"context"
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Updated time.Time
	UserID  int
}

type SnippetModel struct {
	DB *sql.DB
}

// Get returns a snippet from db, given its id.
//
// If snippet does not exist, ErrNoRecord will be returned.
func (m *SnippetModel) Get(ctx context.Context, id int) (Snippet, error) {
	var s Snippet

	row := m.DB.QueryRowContext(ctx, "SELECT * FROM snippets WHERE id=$1", id)

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Updated, &s.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return s, ErrNoRecord
		}
		return s, err
	}

	return s, nil
}

// Update updates a single snippet record in the db.
func (m *SnippetModel) Update(ctx context.Context, user User) error {

	return nil
}

// Insert inserts new snippet into the db.
func (m *SnippetModel) Insert(ctx context.Context, title, content string, userID int) error {

	return nil
}

// Delete deletes a snippet from db.
func (m *SnippetModel) Delete(ctx context.Context, id int) error {
	_, err := m.DB.ExecContext(ctx, "DELETE FROM snippets WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// Latest returns n latest snippets from db.
func (m *SnippetModel) Latest(ctx context.Context, n int) ([]Snippet, error) {

	return nil, nil
}
