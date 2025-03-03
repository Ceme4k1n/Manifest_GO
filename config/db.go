package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func ConnectDB() {
	_ = godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL != "" {
		log.Fatal("❌ DB_URL не найден в .env")
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatal("❌ Ошибка парсинга DB_URL:", err)
	}

	config.MaxConns = 2
	config.HealthCheckPeriod = 30 * time.Second

	DB, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("❌ Не удалось подключиться к БД:", err)
	}
	log.Println("✅ Подключение к БД успешно")
}
