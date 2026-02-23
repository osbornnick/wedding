// Package services contains the business logic and data-access layer.
// GuestService provides CRUD operations for the guests table.
package services

import (
	"context"
	"fmt"
	"log"
	"sort"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lithammer/fuzzysearch/fuzzy"

	"wedding/services/models"
)

// GuestService wraps the connection pool and exposes guest-related operations.
type GuestService struct {
	db *pgxpool.Pool
}

// NewGuestService constructs a GuestService with the given connection pool.
func NewGuestService(db *pgxpool.Pool) *GuestService {
	return &GuestService{db: db}
}

// GetAll returns every guest record from the database.
func (s *GuestService) GetAll() ([]models.Guest, error) {
	rows, err := s.db.Query(context.Background(), `SELECT id, name, invitation_id, aliases, created_at FROM guests ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("GuestService.GetAll query: %w", err)
	}
	defer rows.Close()

	var guests []models.Guest
	for rows.Next() {
		var g models.Guest
		if err := rows.Scan(&g.ID, &g.Name, &g.InvitationID, &g.Aliases, &g.CreatedAt); err != nil {
			return nil, fmt.Errorf("GuestService.GetAll scan: %w", err)
		}
		guests = append(guests, g)
	}
	return guests, rows.Err()
}

// GetByID returns the guest with the given id, or pgx.ErrNoRows if not found.
func (s *GuestService) GetByID(id int) (*models.Guest, error) {
	var g models.Guest
	err := s.db.QueryRow(
		context.Background(),
		`SELECT id, name, invitation_id, aliases, created_at FROM guests WHERE id = $1`, id,
	).Scan(&g.ID, &g.Name, &g.InvitationID, &g.Aliases, &g.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("GuestService.GetByID: %w", err)
	}
	return &g, nil
}

// Create inserts a new guest and returns the record with its generated id.
func (s *GuestService) Create(g *models.Guest) (*models.Guest, error) {
	err := s.db.QueryRow(
		context.Background(),
		`INSERT INTO guests (name, invitation_id, aliases) VALUES ($1, $2, $3) RETURNING id`,
		g.Name, g.InvitationID, g.Aliases,
	).Scan(&g.ID)
	if err != nil {
		return nil, fmt.Errorf("GuestService.Create: %w", err)
	}
	return g, nil
}

// Update overwrites the name and invitation_id for the guest with the given id.
func (s *GuestService) Update(id int, g *models.Guest) (*models.Guest, error) {
	tag, err := s.db.Exec(
		context.Background(),
		`UPDATE guests SET name = $1, invitation_id = $2 WHERE id = $3`,
		g.Name, g.InvitationID, id,
	)
	if err != nil {
		return nil, fmt.Errorf("GuestService.Update: %w", err)
	}
	// Confirm that exactly one row was affected.
	if tag.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}
	g.ID = id
	return g, nil
}

// Delete removes the guest with the given id.
func (s *GuestService) Delete(id int) error {
	tag, err := s.db.Exec(context.Background(), `DELETE FROM guests WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("GuestService.Delete: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (s *GuestService) FuzzySearchByName(name string) ([]models.Guest, error) {
	all, err := s.GetAll()
	if err != nil {
		return nil, fmt.Errorf("GuestService.FuzzySearchByName: %w", err)
	}
	words := []string{}
	guestMap := map[string]*models.Guest{}
	for _, g := range all {
		words = append(words, g.Name)
		guestMap[g.Name] = &g
		for _, a := range g.Aliases {
			words = append(words, a)
			guestMap[a] = &g
		}
	}
	matches := fuzzy.RankFindNormalizedFold(name, words)
	results := []models.Guest{}
	sort.Sort(matches)
	log.Printf("matches for fuzzy search: %v", matches)
	log.Printf("guest map: %v", guestMap)
	for _, match := range matches {
		log.Printf("checking match: %v", match)
		if g, ok := guestMap[match.Target]; ok {
			results = append(results, *g)
		}
	}
	return results, nil
}
