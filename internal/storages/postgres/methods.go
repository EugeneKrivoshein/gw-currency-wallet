package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(provider *PostgresProvider) *PostgresStorage {
	return &PostgresStorage{db: provider.db}
}

func (ps *PostgresStorage) RegisterUser(username, password, email string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`
	_, err = ps.db.Exec(query, username, hash, email)
	return err
}

func (ps *PostgresStorage) AuthenticateUser(username, password string) (bool, error) {
	var storedHash string

	query := `SELECT password FROM users WHERE username = $1`
	err := ps.db.QueryRow(query, username).Scan(&storedHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil // Если пользователь не найден, возвращаем false
		}
		return false, fmt.Errorf("ошибка при аутентификации: %w", err)
	}

	// Сравниваем хеш пароля
	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)); err != nil {
		return false, nil
	}

	return true, nil
}
