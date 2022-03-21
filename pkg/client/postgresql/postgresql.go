package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/to77e/sudoku-api/internal/config"
	"github.com/to77e/sudoku-api/pkg/logging"
)

func NewClient(ctx context.Context, maxAttempts int, sc config.StorageConfig) (conn *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)
	logger := logging.GetLogger()
	logger.Infof("dsn: %v", dsn)
	conn, err = pgxpool.Connect(ctx, dsn)
	return
}
