CREATE TABLE IF NOT EXISTS bookings (
    id                  INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id             INT         NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    offer_id            INT         NOT NULL REFERENCES offers(id) ON DELETE CASCADE,
    status              VARCHAR(30) NOT NULL DEFAULT 'reserved',
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

CREATE INDEX IF NOT EXISTS idx_bookings_user_id ON bookings(user_id);
CREATE INDEX IF NOT EXISTS idx_bookings_offer_id ON bookings(offer_id);
CREATE INDEX IF NOT EXISTS idx_bookings_status   ON bookings(status);