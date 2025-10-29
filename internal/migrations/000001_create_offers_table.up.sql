CREATE TABLE IF NOT EXISTS offers (
    id                  INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    flight_no           VARCHAR(20) NOT NULL,
    airline             VARCHAR(100) NOT NULL,
    departure_airport   CHAR(3) NOT NULL,
    arrival_airport     CHAR(3) NOT NULL,
    departure_at        TIMESTAMPTZ NOT NULL,
    arrival_at          TIMESTAMPTZ NOT NULL,
    currency            CHAR(3) NOT NULL DEFAULT 'UZS',
    price_tiyin          BIGINT NOT NULL CHECK (price_tiyin >= 0),
    seats_left          INT NOT NULL CHECK (seats_left >= 0),
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
    CONSTRAINT chk_offers_time CHECK (arrival_at > departure_at)          -- прилёт > вылет
    );

CREATE INDEX IF NOT EXISTS idx_offers_route_date
    ON offers(departure_airport, arrival_airport, departure_at);