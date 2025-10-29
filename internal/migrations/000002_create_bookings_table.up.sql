CREATE TABLE IF NOT EXISTS bookings (
    id                  INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    offer_id            INT         NOT NULL REFERENCES offers(id) ON DELETE CASCADE,
    passenger_name      VARCHAR(50) NOT NULL,
    passenger_surname   VARCHAR(50) NOT NULL,
    passenger_birth     DATE        NOT NULL,
    passport_doc        VARCHAR(30) NOT NULL,
    status              VARCHAR(30) NOT NULL DEFAULT 'booked',
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

CREATE INDEX IF NOT EXISTS idx_bookings_offer_id ON bookings(offer_id);
CREATE INDEX IF NOT EXISTS idx_bookings_status   ON bookings(status);