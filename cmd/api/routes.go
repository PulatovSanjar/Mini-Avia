package main

import (
	"Mini-Avia/internal/bookings"
	"Mini-Avia/internal/offers"
	"Mini-Avia/internal/tickets"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net/http"
)

func loadRoutes(mux *http.ServeMux, pool *pgxpool.Pool, log *slog.Logger) {
	off := offers.NewHandler(pool, log)
	bok := bookings.NewHandler(pool, log)
	tkt := tickets.NewHandler(pool, log)

	//mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("pong")) })

	mux.HandleFunc("GET /offers", off.Search)
	mux.HandleFunc("POST /bookings", bok.Create)
	mux.HandleFunc("POST /tickets/{id}/issue", tkt.Issue)
}
