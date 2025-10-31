package middleware

import (
	"Mini-Avia/cmd/api/middleware"
	"Mini-Avia/internal/common"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := common.MustLoad()
	log := common.NewLogger(cfg.LogLevel)
	log.Info("Starting Mini-Avia", "port", cfg.Port)

	ctx := context.Background()
	pool, err := common.NewDB(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	// (опционально) fail-fast ping БД
	pctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	if err := pool.Ping(pctx); err != nil {
		cancel()
		log.Error("DB ping failed", "error", err)
		os.Exit(1)
	}
	cancel()

	router := http.NewServeMux()
	loadRoutes(router)

	stack := middleware.CreateStack()

	// собираем маршруты и оборачиваем middleware-логгером
	handler := requestLogger(log, routes(pool))

	app := &application{port: cfg.Port}
	go func() {
		if err := app.serve(handler); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	// graceful shutdown
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

type wrapWriter struct {
	http.ResponseWriter
	status int
}

func (w *wrapWriter) WriteHeader(code int) { w.status = code; w.ResponseWriter.WriteHeader(code) }

func requestLogger(log *slog.Logger, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := &wrapWriter{ResponseWriter: w, status: http.StatusOK}
		handler.ServeHTTP(ww, r)
		log.Info("http",
			"method", r.Method,
			"path", r.URL.Path,
			"status", ww.status,
			"duration_ms", time.Since(start).Milliseconds(),
			"remote", r.RemoteAddr,
		)
	})
}
