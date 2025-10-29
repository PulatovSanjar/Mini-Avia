package main

//
//import (
//	"Mini-Avia/internal/bookings"
//	"context"
//	"log/slog"
//	"net/http"
//
//	"github.com/jackc/pgx/v5/pgxpool"
//)
//
//// only ограничивает метод (аналог Allow в фреймворках)
//
//
////// пример "param" роутинга без внешних либ (для /bookings/{id})
////func withBookingID(base string, h http.Handler) http.Handler {
////	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
////		// /api/v1/bookings/<id>
////		if !strings.HasPrefix(r.URL.Path, base+"/bookings/") {
////			http.NotFound(w, r)
////			return
////		}
////		id := strings.TrimPrefix(r.URL.Path, base+"/bookings/")
////		if id == "" || strings.Contains(id, "/") {
////			http.NotFound(w, r)
////			return
////		}
////		ctx := context.WithValue(r.Context(), bookings.CtxKeyID{}, id)
////		h.ServeHTTP(w, r.WithContext(ctx))
////	})
////}
//
//func routes(pool *pgxpool.Pool, log *slog.Logger) http.Handler {
//	handler := &bookings.Handler{}
//	router := http.NewServeMux()
//	//base := "/api/v1"
//
//	router.HandleFunc("POST /bookings", handler.findById)
//
//	//// GET /api/v1/offers?from=...&to=...&date=YYYY-MM-DD
//	router.HandleFunc(base+"/offers", api.only(http.MethodGet,
//	//	offers.ListHandler(pool, log)))
//	//
//	//// POST /api/v1/bookings
//	//mux.Handle(base+"/bookings", only(http.MethodPost,
//	//	bookings.CreateHandler(pool, log)))
//	//
//	//// GET /api/v1/bookings/{id} (пример параметра path)
//	//mux.Handle(base+"/bookings/", only(http.MethodGet,
//	//	withBookingID(base, bookings.GetByIDHandler(pool, log))))
//	//
//	//// POST /api/v1/tickets/issue
//	//mux.Handle(base+"/tickets/issue", only(http.MethodPost,
//	//	tickets.IssueHandler(pool, log)))
//
//	// можно добавить /ping без БД, если нужно
//	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("pong")) })
//
//	return mux
//}
