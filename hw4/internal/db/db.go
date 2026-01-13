package db

import(
	"context"
	"errors"
	"os"
	"path/filepath"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func NewPostgresDB(ctx context.Context) (*pgxpool.Pool, error){
	// Try to load .env from current directory or parent directories
	_ = godotenv.Load(".env")
	_ = godotenv.Load("../.env")
	_ = godotenv.Load(filepath.Join("..", "..", ".env"))

	connStr := os.Getenv("POSTGRES_CONN_STR")
	if connStr == "" {
		return nil, errors.New("POSTGRES_CONN_STR environment variable is not set")
	}
	
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return pool, nil
}