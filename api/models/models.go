// Package models defines the data structures that mirror the database schema.
// Each struct maps directly to a table defined in schema/sql/.
package models

import "time"

// ── Guest ────────────────────────────────────────────────────────────────────

// Guest represents a single wedding guest (schema: guests table).
type Guest struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	InvitationID int       `json:"invitation_id"`
	Aliases      []string  `json:"aliases"`
	CreatedAt    time.Time `json:"created_at"`
}

// ── Invitation ───────────────────────────────────────────────────────────────

// Invitation represents a mailed invitation sent to a household
// (schema: invitations table).
type Invitation struct {
	ID        int       `json:"id"`
	Address   string    `json:"address"`
	NumGuests int       `json:"num_guests"`
	CreatedAt time.Time `json:"created_at"`
}

// ── Response ─────────────────────────────────────────────────────────────────

// Response captures an RSVP reply from an invitation recipient
// (schema: responses table).
type Response struct {
	ID                    int       `json:"id"`
	AttendingWedding      bool      `json:"attending_wedding"`
	AttendingWeddingCount int       `json:"attending_wedding_count"`
	AttendingFriday       bool      `json:"attending_friday"`
	AttendingFridayCount  int       `json:"attending_friday_count"`
	AttendingBrunch       bool      `json:"attending_brunch"`
	AttendingBrunchCount  int       `json:"attending_brunch_count"`
	DietaryRestrictions   string    `json:"dietary_restrictions"`
	InvitationID          int       `json:"invitation_id"`
	CreatedAt             time.Time `json:"created_at"`
}

// ── Gift ─────────────────────────────────────────────────────────────────────

// Gift represents an item on the wedding gift registry
// (schema: gifts table).
type Gift struct {
	ID          int       `json:"id"`
	Img         string    `json:"img"`
	Name        string    `json:"name"`
	Link        string    `json:"link"`
	Progress    float64   `json:"progress"`
	Total       float64   `json:"total"`
	Purchased   bool      `json:"purchased"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
