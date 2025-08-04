package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func openTestDB(t *testing.T) (*sql.DB, error) {
	dsn := os.Getenv("MARKBIN_DB_DSN_TEST")
	if dsn == "" {
		return nil, fmt.Errorf("dsn can not be empty")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	m, err := migrate.New("file://../../migrations", dsn)
	if err != nil {
		return nil, err
	}

	if err = m.Up(); err != nil {
		return nil, err
	}

	insertTestData(db)

	t.Cleanup(cleanupTestDB)

	return db, nil
}

func insertTestData(db *sql.DB) {
	f, err := os.Open("./testdata/testdata.sql")
	if err != nil {
		log.Fatalf("failed to insert test data into db: %v", err)
	}

	buff := new(bytes.Buffer)
	_, err = io.Copy(buff, f)
	if err != nil {
		log.Fatalf("failed to insert test data into db: %v", err)
	}

	stmns := buff.String()

	_, err = db.Exec(stmns)
	if err != nil {
		log.Fatalf("failed to insert test data into db: %v", err)
	}

}

func cleanupTestDB() {
	dsn := os.Getenv("MARKBIN_DB_DSN_TEST")
	if dsn == "" {
		log.Fatalf("failed to cleanup db: dsn can not be empty")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to cleanup db: %v", err)
	}
	defer db.Close()

	f, err := os.Open("./testdata/teardown.sql")
	if err != nil {
		log.Fatalf("failed to cleanup db: %v", err)
	}

	buff := new(bytes.Buffer)
	_, err = io.Copy(buff, f)
	if err != nil {
		log.Fatalf("failed to cleanup db: %v", err)
	}

	stmns := buff.String()

	_, err = db.Exec(stmns)
	if err != nil {
		log.Fatalf("failed to cleanup db: %v", err)
	}

}
