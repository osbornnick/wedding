-- V6__init_purchases.sql
-- Migration: create purchases table

CREATE TABLE IF NOT EXISTS purchases (
    id SERIAL PRIMARY KEY,
    gift_id INT NOT NULL REFERENCES gifts(id) ON DELETE CASCADE,
    guest_id INT NOT NULL REFERENCES guests(id) ON DELETE CASCADE,
    amount NUMERIC(10,2),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
