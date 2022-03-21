package sudoku

import (
	"context"

	"github.com/to77e/sudoku-api/internal/models"
)

type Repository interface {
	Create(ctx context.Context, s *models.Sudoku) error
	FindOne(ctx context.Context, id string) (*models.Sudoku, error)
}
