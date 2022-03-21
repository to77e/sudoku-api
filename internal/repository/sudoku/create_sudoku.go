package sudoku

import (
	"context"
	"fmt"

	"github.com/lib/pq"
	"github.com/to77e/sudoku-api/internal/models"
)

func (r *Repository) Create(ctx context.Context, s *models.Sudoku) (err error) {
	const query = `
		insert into public.sudoku (sudoku, k)
		values ($1, $2)
		returning uuid;
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(query)))
	err = r.client.QueryRow(ctx, query, pq.Array(s.Field), s.K).Scan(&s.UUID)
	return
}
