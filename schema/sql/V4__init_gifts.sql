-- V4__init_gifts.sql
-- Migration: create gifts table

CREATE TABLE IF NOT EXISTS gifts (
    id SERIAL PRIMARY KEY,
    img VARCHAR(512),
    name VARCHAR(255) NOT NULL,
    link TEXT,
    progress NUMERIC(10,2) DEFAULT 0.00,
    total NUMERIC(10,2) DEFAULT 0.00,
    purchased BOOLEAN DEFAULT FALSE,
    description TEXT
);
