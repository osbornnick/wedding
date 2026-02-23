-- V1__init_guests.sql
-- Migration: create guests table

CREATE TABLE IF NOT EXISTS guests (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    invitation_id INT,
    aliases VARCHAR(255)[],
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
