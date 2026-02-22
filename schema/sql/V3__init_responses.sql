-- V3__init_responses.sql
-- Migration: create responses table

CREATE TABLE IF NOT EXISTS responses (
    id SERIAL PRIMARY KEY,
    attending_wedding BOOLEAN NOT NULL DEFAULT FALSE,
    attending_wedding_count INTEGER NOT NULL DEFAULT 0,
    attending_friday BOOLEAN NOT NULL DEFAULT FALSE,
    attending_friday_count INTEGER NOT NULL DEFAULT 0,
    attending_brunch BOOLEAN NOT NULL DEFAULT FALSE,
    attending_brunch_count INTEGER NOT NULL DEFAULT 0,
    dietary_restrictions TEXT
);
