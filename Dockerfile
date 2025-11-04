# -------- deps (опциональный кеш, оставим как прогрев) --------
FROM golang:1.25.3 AS deps
WORKDIR /app
COPY go.mod ./
RUN go mod download

# -------- DEV (Air + hot reload) --------
FROM golang:1.25.3 AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air", "-c", ".air.toml"]

# -------- PROD build --------
#FROM golang:1.25.3 AS builder
#WORKDIR /app
#COPY go.mod ./
#RUN go mod download
#COPY . ./
#RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o /api ./cmd/api
#
## -------- PROD runtime --------
#FROM gcr.io/distroless/base-debian12 AS prod
#WORKDIR /srv
#COPY --from=builder --chmod=0755 --chown=nonroot:nonroot /api /bin/api
#EXPOSE 8080
#USER nonroot:nonroot
#ENTRYPOINT ["/bin/api"]
