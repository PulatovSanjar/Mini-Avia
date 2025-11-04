package tickets

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type Handler struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewHandler(db *pgxpool.Pool, log *slog.Logger) *Handler { return &Handler{db: db, log: log} }
