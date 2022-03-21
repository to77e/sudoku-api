package sudoku

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/to77e/sudoku-api/pkg/logging"
)

type Repository struct {
	client *pgxpool.Pool
	logger *logging.Logger
}

func NewRepository(client *pgxpool.Pool, logger *logging.Logger) *Repository {
	return &Repository{
		client: client,
		logger: logger,
	}
}
