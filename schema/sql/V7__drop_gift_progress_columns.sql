-- V7__drop_gift_progress_columns.sql
-- Migration: remove progress, total, and purchased from gifts
-- These values are now derived from the purchases table.

ALTER TABLE gifts
    DROP COLUMN IF EXISTS progress,
    DROP COLUMN IF EXISTS total,
    DROP COLUMN IF EXISTS purchased;
