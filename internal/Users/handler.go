package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	DB     *pgxpool.Pool
	Secret []byte
}

func NewHandler(db *pgxpool.Pool, secret string) *Handler {
	return &Handler{DB: db, Secret: []byte(secret)}
}

type registerReq struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	BirthDate   string `json:"birth_date"` // YYYY-MM-DD
	PassportDoc string `json:"passport_doc"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type tokenResp struct {
	Token string `json:"token"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerReq
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	email := strings.ToLower(strings.TrimSpace(req.Email))
	if email == "" || len(req.Password) < 8 {
		http.Error(w, "invalid email or password", http.StatusBadRequest)
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	var id int
	err := h.DB.QueryRow(ctx, `
		INSERT INTO users (name, surname, birth_date, passport_doc, email, password_hash)
		VALUES ($1,$2,$3,$4,$5,$6) RETURNING id
	`, req.Name, req.Surname, req.BirthDate, req.PassportDoc, email, string(hash)).Scan(&id)
	if err != nil {
		http.Error(w, "email already exist", http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(tokenResp{Token: h.makeJWT(id)})
	if err != nil {
		return
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginReq
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	email := strings.ToLower(strings.TrimSpace(req.Email))

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	var id int
	var hash string
	err := h.DB.QueryRow(ctx, `SELECT id, password_hash FROM users WHERE email=$1`, email).Scan(&id, &hash)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password)) != nil {
		http.Error(w, "invalid credentials", 401)
		return
	}
	err = json.NewEncoder(w).Encode(tokenResp{Token: h.makeJWT(id)})
	if err != nil {
		return
	}
}

func (h *Handler) makeJWT(id int) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	s, _ := t.SignedString(h.Secret)
	return s
}
