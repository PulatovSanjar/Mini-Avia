// internal/bookings/create_mock_test.go
package bookings

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"Mini-Avia/internal/middleware"

	"github.com/golang-jwt/jwt/v5"
	pgxmock "github.com/pashagolub/pgxmock/v3"
	"log/slog"
)

func TestCreate_HappyPath_WithPgxMock(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(mock.Close)

	// хендлер + auth middleware (положит user_id=42 в контекст)
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	h := NewHandler(mock, log)
	secret := []byte("test-secret")
	handler := middleware.RequireAuth(secret)(http.HandlerFunc(h.Create))

	// ожидаемые SQL в транзакции
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE offers SET seats_left = seats_left - 1\s+WHERE id = \$1 AND seats_left > 0`).
		WithArgs(int64(777)).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	rows := pgxmock.NewRows([]string{"id", "offer_id", "user_id", "status", "created_at"}).
		AddRow(int64(1001), int64(777), int64(42), "reserved", time.Now())

	mock.ExpectQuery(`INSERT INTO bookings`).
		WithArgs(int64(777), 42, "reserved").
		WillReturnRows(rows)

	mock.ExpectCommit()

	// HTTP-запрос с JWT
	token := mustJWT(t, 42, secret)
	req := httptest.NewRequest(http.MethodPost, "/bookings", strings.NewReader(`{"offer_id":777}`))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// вызов
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rr.Code, rr.Body.String())
	}
	var resp Booking
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.OfferID != 777 || resp.UserID != 42 || resp.Status != "reserved" {
		t.Fatalf("unexpected resp: %+v", resp)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestCreate_NoSeats_WithPgxMock(t *testing.T) { // этот тест будет ломаться чтобы понять как ломаются тесты
	mock, _ := pgxmock.NewPool()
	t.Cleanup(mock.Close)

	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	h := NewHandler(mock, log)
	secret := []byte("test-secret")
	handler := middleware.RequireAuth(secret)(http.HandlerFunc(h.Create))

	// UPDATE не изменил ни одной строки → err "no_seats" → 409
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE offers SET seats_left = seats_left - 1\s+WHERE id = \$1 AND seats_left > 0`).
		WithArgs(int64(555)).
		WillReturnResult(pgxmock.NewResult("UPDATE", 0))
	// при ошибке транзакция откатится
	mock.ExpectRollback()

	token := mustJWT(t, 42, secret)
	req := httptest.NewRequest(http.MethodPost, "/bookings", strings.NewReader(`{"offer_id":555}`))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code == http.StatusConflict {
		t.Fatalf("status=%d body=%s (want 409)", rr.Code, rr.Body.String())
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func mustJWT(t *testing.T, uid int64, secret []byte) string {
	t.Helper()
	claims := jwt.MapClaims{
		"sub": uid,
		"exp": time.Now().Add(10 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	}
	s, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		t.Fatalf("sign: %v", err)
	}
	return s
}
