package tickets

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewHandler(db *pgxpool.Pool, log *slog.Logger) *Handler { return &Handler{db: db, log: log} }

type Ticket struct {
	ID        int64     `json:"id"`
	BookingID int64     `json:"booking_id"`
	Number    string    `json:"ticket_number"`
	IssuedAt  time.Time `json:"issued_at"`
}

func (h *Handler) Issue(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "id required", http.StatusBadRequest)
		return
	}
	bookingID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || bookingID <= 0 {
		http.Error(w, "id must be positive integer", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	var t Ticket
	err = h.withTx(ctx, func(tx pgx.Tx) error {
		var status string
		if err := tx.QueryRow(ctx,
			`SELECT status FROM bookings WHERE id = $1 FOR UPDATE`,
			bookingID,
		).Scan(&status); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return errors.New("not_found")
			}
			return err
		}
		if status != "reserved" {
			return errors.New("bad_status")
		}

		if err := tx.QueryRow(ctx, `
			INSERT INTO tickets (booking_id, ticket_number)
			VALUES ($1, 'TKT-' || EXTRACT(EPOCH FROM clock_timestamp())::bigint)
			RETURNING id, booking_id, ticket_number, issued_at
		`, bookingID).Scan(&t.ID, &t.BookingID, &t.Number, &t.IssuedAt); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				return errors.New("already_issued")
			}
			return err
		}

		if _, err := tx.Exec(ctx,
			`UPDATE bookings SET status = 'issued' WHERE id = $1`,
			bookingID,
		); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		switch err.Error() {
		case "not_found":
			http.Error(w, "booking not found", http.StatusNotFound)
			return
		case "bad_status":
			http.Error(w, "booking not in reserved state", http.StatusConflict)
			return
		case "already_issued":
			http.Error(w, "ticket already issued", http.StatusConflict)
			return
		default:
			h.log.Error("ticket_issue_failed", "err", err)
			http.Error(w, "internal", http.StatusInternalServerError)
			return
		}
	}

	h.log.Info("ticket_issued", "booking_id", t.BookingID, "ticket", t.Number)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(t)
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
