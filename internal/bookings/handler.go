package bookings

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net/http"
	"time"
)

type Handler struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewHandler(db *pgxpool.Pool, log *slog.Logger) *Handler { return &Handler{db: db, log: log} }

type CreateRequest struct {
	OfferID int64  `json:"offer_id"`
	UserID  string `json:"user_id"`
}

type Booking struct {
	ID      int64     `json:"id"`
	OfferID int64     `json:"offer_id"`
	UserID  string    `json:"user_id"`
	Status  string    `json:"status"`
	Created time.Time `json:"created_at"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil ||
		req.OfferID <= 0 ||
		req.UserID == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	var booking Booking
	err := h.withTx(ctx, func(tx pgx.Tx) error {
		ct, err := tx.Exec(ctx, `
			UPDATE offers SET seats_left = seats_left - 1
			WHERE id = $1 AND seats_left > 0
		`, req.OfferID)

		if err != nil {
			return err
		}
		if ct.RowsAffected() != 1 {
			return errors.New("no_seats")
		}

		return tx.QueryRow(ctx, `
            INSERT INTO bookings (
                offer_id, user_id, status
            ) VALUES ($1, $2, $3)
            RETURNING id, offer_id, user_id, status, created_at
        `,
			req.OfferID, req.UserID, "reserved",
		).Scan(
			&booking.ID, &booking.OfferID, &booking.UserID, &booking.Status, &booking.Created,
		)
	})

	if err != nil {
		if err.Error() == "no_seats" {
			http.Error(w, "no seats", http.StatusConflict)
			return
		}
		h.log.Error("booking_create_failed", "err", err)
		http.Error(w, "booking_create_failed", http.StatusInternalServerError)
		return
	}

	h.log.Info("booking_created", "booking_id", booking.ID, "offer_id", booking.OfferID)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(booking)
}

func (h *Handler) withTx(ctx context.Context, fn func(pgx.Tx) error) error {
	tx, err := h.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if err := fn(tx); err != nil {
		return err
	}
	return tx.Commit(ctx)
}
