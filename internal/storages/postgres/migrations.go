package postgres

import (
	"database/sql"
	"log"
)

func RunMigrations(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) NOT NULL UNIQUE,
        password TEXT NOT NULL,
		email VARCHAR(255) UNIQUE
    );
    `
	if _, err := db.Exec(query); err != nil {
		return err
	}

	log.Println("Миграции выполнены успешно.")
	return nil
}
