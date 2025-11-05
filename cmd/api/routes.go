package main

import (
	auth "Mini-Avia/internal/Users"
	"Mini-Avia/internal/bookings"
	"Mini-Avia/internal/middleware"
	"Mini-Avia/internal/offers"
	"Mini-Avia/internal/tickets"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net/http"
	"os"
)

func loadRoutes(mux *http.ServeMux, pool *pgxpool.Pool, log *slog.Logger) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	auh := auth.NewHandler(pool, os.Getenv("JWT_SECRET"))
	reqAuth := middleware.RequireAuth(secret)
	off := offers.NewHandler(pool, log)
	bok := bookings.NewHandler(pool, log)
	tkt := tickets.NewHandler(pool, log)

	mux.HandleFunc("POST /auth/register", auh.Register)
	mux.HandleFunc("POST /auth/login", auh.Login)

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			return
		}
	})

	mux.HandleFunc("GET /all-offers", off.GetAll)
	mux.HandleFunc("GET /offers", off.Search)

	mux.Handle("POST /bookings", reqAuth(http.HandlerFunc(bok.Create)))
	mux.Handle("POST /tickets/{id}/issue", reqAuth(http.HandlerFunc(tkt.Issue)))
}
