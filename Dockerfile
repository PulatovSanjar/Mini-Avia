# -------- build stage --------
FROM golang:1.25.3 AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . ./

RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o /api ./cmd/api


# -------- runtime stage --------
FROM gcr.io/distroless/base-debian12
WORKDIR /srv
COPY --from=builder --chmod=0755 --chown=nonroot:nonroot /api /bin/api
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/bin/api"]
