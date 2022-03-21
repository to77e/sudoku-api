package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/to77e/sudoku-api/internal/config"
	"github.com/to77e/sudoku-api/internal/service/sudoku"
	"github.com/to77e/sudoku-api/pkg/client/postgresql"
	"github.com/to77e/sudoku-api/pkg/logging"
)

const maxAttempts = 3

func main() {
	logger := logging.GetLogger()
	ctx := context.Background()
	cfg := config.GetConfig()

	logger.Info("create router")
	router := httprouter.New()

	logger.Info("create pg client")
	pgClient, err := postgresql.NewClient(ctx, maxAttempts, cfg.Storage)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("create new sudoku service")
	sudokuService := sudoku.NewSudoku(pgClient, logger)

	logger.Info("register sudoku handler")
	handler := sudoku.NewHanlder(sudokuService, logger)
	handler.Register(router)

	logger.Info("listen to tcp")
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("server is listening to port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	logger.Fatalln(server.Serve(listener))
}
