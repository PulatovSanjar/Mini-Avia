package middleware

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net/http"
)

//
//func routes() http.Handler {
//	mux := http.NewServeMux()
//	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
//
//		w.Write([]byte("pong"))
//	})
//	return mux
//}

// import (
//
//	"Mini-Avia/internal/bookings"
//	"context"
//	"log/slog"
//	"net/http"
//
//	"github.com/jackc/pgx/v5/pgxpool"
//
// )
//
// // only ограничивает метод (аналог Allow в фреймворках)
//
// //// пример "param" роутинга без внешних либ (для /bookings/{id})
// //func withBookingID(base string, h http.Handler) http.Handler {
// //	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// //		// /api/v1/bookings/<id>
// //		if !strings.HasPrefix(r.URL.Path, base+"/bookings/") {
// //			http.NotFound(w, r)
// //			return
// //		}
// //		id := strings.TrimPrefix(r.URL.Path, base+"/bookings/")
// //		if id == "" || strings.Contains(id, "/") {
// //			http.NotFound(w, r)
// //			return
// //		}
// //		ctx := context.WithValue(r.Context(), bookings.CtxKeyID{}, id)
// //		h.ServeHTTP(w, r.WithContext(ctx))
// //	})
// //}
func loadRoutes(pool *pgxpool.Pool) http.Handler {
	handler := &bookings.Handler{}
	router := http.NewServeMux()
	//base := "/api/v1"

	router.HandleFunc("POST /bookings", handler.findById)

	// можно добавить /ping без БД, если нужно
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("pong")) })

	return router
}
