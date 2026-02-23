-- V2__init_invitations.sql
-- Migration: create invitations table

CREATE TABLE IF NOT EXISTS invitations (
    id SERIAL PRIMARY KEY,
    address VARCHAR(255),
    num_guests INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
