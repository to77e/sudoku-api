package sudoku

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/to77e/sudoku-api/internal/apperror"
	"github.com/to77e/sudoku-api/internal/handlers"
	"github.com/to77e/sudoku-api/internal/models"
	"github.com/to77e/sudoku-api/pkg/logging"
)

type handler struct {
	sudoku *Service
	logger *logging.Logger
}

func NewHanlder(s *Service, logger *logging.Logger) handlers.Handler {
	return &handler{
		sudoku: s,
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, "/sudoku/:k", apperror.Middleware(h.CreateSudoku))
	router.HandlerFunc(http.MethodGet, "/sudoku/:uuid", apperror.Middleware(h.GetSudoku))
}

func (h *handler) CreateSudoku(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()

	params := httprouter.ParamsFromContext(r.Context())
	k, err := strconv.Atoi(params.ByName("k"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	start := time.Now()
	s := &models.Sudoku{
		K:     k,
		Field: GenerateField(k),
	}
	h.logger.Infof("field generated in %v\n", time.Since(start))

	if err := h.sudoku.repository.Create(ctx, s); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data, _ := json.Marshal(s)
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) GetSudoku(w http.ResponseWriter, r *http.Request) error {
	// this handler should get sudoku by id and check it
	return apperror.ErrNotFound
}
