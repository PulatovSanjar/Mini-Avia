package offers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewHandler(db *pgxpool.Pool, log *slog.Logger) *Handler { return &Handler{db: db, log: log} }

type Offer struct {
	ID         int64     `json:"id"`
	FlightNo   string    `json:"flight_no"`
	Airline    string    `json:"airline"`
	DesAirport string    `json:"departure_airport"`
	ArrAirport string    `json:"arrival_airport"`
	DepTime    time.Time `json:"departure_at"`
	ArrTime    time.Time `json:"arrival_at"`
	Currency   string    `json:"currency"`
	Price      int64     `json:"price_tiyin"`
	SeatsLeft  int       `json:"seats_left"`
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	q := r.URL.Query()
	from, to, date := q.Get("from"), q.Get("to"), q.Get("date")
	if len(from) != 3 || len(to) != 3 || date == "" {
		http.Error(w, "params: from(3), to(3), date(YYYY-MM-DD)", http.StatusBadRequest)
		return
	}

	h.log.Info("offers_query", "from", from, "to", to, "date", date)

	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, date)

	rows, err := h.db.Query(ctx, `
		SELECT id, flight_no, airline, departure_airport, arrival_airport,
		       departure_at, arrival_at, currency, price_tiyin, seats_left
		FROM offers
		WHERE departure_airport = $1 AND arrival_airport = $2 AND DATE(departure_at) = $3
		ORDER BY price_tiyin
	`, from, to, parsedDate)

	if err != nil {
		h.log.Error("offers_query_failed", "err", err)
		http.Error(w, "sql error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var res []Offer
	for rows.Next() {
		var o Offer
		if err := rows.Scan(&o.ID, &o.FlightNo, &o.Airline, &o.DesAirport, &o.ArrAirport, &o.DepTime, &o.ArrTime, &o.Currency, &o.Price, &o.SeatsLeft); err != nil {
			h.log.Error("offers_scan_failed", "err", err)
			http.Error(w, "scan error", http.StatusInternalServerError)
			return
		}
		res = append(res, o)
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	nextDay := time.Now().Add(24 * time.Hour).Truncate(24 * time.Hour)

	rows, err := h.db.Query(ctx, `
		SELECT id, flight_no, airline, departure_airport, arrival_airport,
		       departure_at, arrival_at, currency, price_tiyin, seats_left
		  FROM offers
		 WHERE departure_at >= $1
		 ORDER BY price_tiyin ASC, departure_at ASC
		 LIMIT 200
	`, nextDay)

	if err != nil {
		h.log.Error("offers_query_failed", "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var res []Offer
	for rows.Next() {
		var o Offer
		if err := rows.Scan(
			&o.ID, &o.FlightNo, &o.Airline, &o.DesAirport, &o.ArrAirport,
			&o.DepTime, &o.ArrTime, &o.Currency, &o.Price, &o.SeatsLeft,
		); err != nil {
			h.log.Error("offers_scan_failed", "err", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		res = append(res, o)
	}
	if err := rows.Err(); err != nil {
		h.log.Error("offers_rows_err", "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(res)
}
