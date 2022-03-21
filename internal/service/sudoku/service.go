package sudoku

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/to77e/sudoku-api/internal/repository/sudoku"
	"github.com/to77e/sudoku-api/pkg/logging"
)

type Service struct {
	repository Repository
}

func NewSudoku(pgClient *pgxpool.Pool, logger *logging.Logger) *Service {
	repository := sudoku.NewRepository(pgClient, logger)
	return &Service{
		repository: repository,
	}
}
