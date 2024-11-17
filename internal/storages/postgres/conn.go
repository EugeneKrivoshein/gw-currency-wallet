package postgres

import (
	"database/sql"
	"fmt"

	"github.com/EugeneKrivoshein/gw-currency-wallet/internal/config"
)

// Структура для работы с базой данных PostgreSQL
type PostgresProvider struct {
	db *sql.DB
}

// Конструктор для создания нового провайдера PostgreSQL
func NewPostgresProvider(cfg *config.Config) (*PostgresProvider, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	return &PostgresProvider{db: db}, nil
}

// Метод для закрытия подключения
func (p *PostgresProvider) Close() error {
	return p.db.Close()
}

// Метод для получения объекта *sql.DB
func (p *PostgresProvider) DB() *sql.DB {
	return p.db
}
