package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Mini-Avia/internal/common"
	"Mini-Avia/internal/middleware"
)

func main() {
	cfg := common.MustLoad()
	log := common.NewLogger(cfg.LogLevel).With("component", "api")
	log.Info("starting", "port", cfg.Port)

	if os.Getenv("JWT_SECRET") == "" {
		err := os.Setenv("JWT_SECRET", "dev-secret-change-me")
		if err != nil {
			return
		}
	}

	ctx := context.Background()
	pool, err := common.NewDB(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Error("db_connect_failed", "err", err)
		os.Exit(1)
	}
	defer pool.Close()

	pctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	if err := pool.Ping(pctx); err != nil {
		cancel()
		log.Error("db_ping_failed", "err", err)
		os.Exit(1)
	}
	cancel()

	router := http.NewServeMux()
	loadRoutes(router, pool, log)

	stack := middleware.CreateStack(
		middleware.Recovery(log),
		middleware.RequestLogger(log),
	)
	handler := stack(router)

	app := &application{port: cfg.Port}
	go func() {
		if err := app.serve(handler); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Info("Shutting down...")

	ctxShutDown, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()
	if err := app.srv.Shutdown(ctxShutDown); err != nil {
		log.Error("shutdown error", "error", err)
	}
	log.Info("Server gracefully stopped")
}
