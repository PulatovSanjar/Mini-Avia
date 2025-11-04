package tickets

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"time"
)

type Handler struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewHandler(db *pgxpool.Pool, log *slog.Logger) *Handler { return &Handler{db: db, log: log} }

type CreateRequest struct {
	OfferID     int64  `json:"offer_id"`
	Name        string `json:"passenger_name"`
	Surname     string `json:"passenger_surname"`
	PassportDoc string `json:"passport_doc"`
	PassBirth   string `json:"passenger_birth"`
}

type Ticket struct {
	ID          int64     `json:"id"`
	OfferID     int64     `json:"offer_id"`
	Status      string    `json:"status"`
	Name        string    `json:"passenger_name"`
	Surname     string    `json:"passenger_surname"`
	PassportDoc string    `json:"passport_doc"`
	PassBirth   string    `json:"email"`
	Created     time.Time `json:"created_at"`
}
