package db

import (
	"database/sql"
	"fmt"
)

func InitTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS organizations (
		id SERIAL PRIMARY KEY,
		org_name TEXT NOT NULL,
		email TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS aid_requests (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		aid_type TEXT NOT NULL,
		status TEXT DEFAULT 'pending',
		org_id INT REFERENCES organizations(id),
		document_url TEXT
	);`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to init tables: %w", err)
	}
	return nil
}
