CREATE TABLE IF NOT EXISTS users (
    id                  INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name                VARCHAR(50) NOT NULL,
    surname             VARCHAR(50) NOT NULL,
    birth_date          DATE        NOT NULL,
    passport_doc        VARCHAR(30) NOT NULL,
    email               VARCHAR(50) NOT NULL UNIQUE,
    password_hash       TEXT        NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );