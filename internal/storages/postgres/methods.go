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

// Конструктор для создания нового хранилища
func NewPostgresStorage(provider *PostgresProvider) *PostgresStorage {
	return &PostgresStorage{db: provider.db}
}

// Пример метода для регистрации пользователя
func (ps *PostgresStorage) RegisterUser(username, password string) error {
	// Генерация хеша пароля
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Вставка нового пользователя в таблицу
	query := `INSERT INTO users (username, password) VALUES ($1, $2)`
	_, err = ps.db.Exec(query, username, hash)
	return err
}

// Аутентификация пользователя
func (ps *PostgresStorage) AuthenticateUser(username, password string) (bool, error) {
	var storedHash string

	// Получаем хеш пароля из базы данных
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
		return false, nil // Неверный пароль
	}

	return true, nil // Успешная аутентификация
}
