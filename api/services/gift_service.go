// GiftService provides CRUD operations for the gifts table.
package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"wedding/services/models"
)

// GiftService wraps the connection pool and exposes gift registry operations.
type GiftService struct {
	db *pgxpool.Pool
}

// NewGiftService constructs a GiftService with the given connection pool.
func NewGiftService(db *pgxpool.Pool) *GiftService {
	return &GiftService{db: db}
}

// GetAll returns every gift record from the database.
func (s *GiftService) GetAll() ([]models.Gift, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT id, img, name, link, progress, total, purchased, description
		FROM gifts ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("GiftService.GetAll query: %w", err)
	}
	defer rows.Close()

	var gifts []models.Gift
	for rows.Next() {
		var g models.Gift
		if err := rows.Scan(
			&g.ID, &g.Img, &g.Name, &g.Link,
			&g.Progress, &g.Total, &g.Purchased, &g.Description,
		); err != nil {
			return nil, fmt.Errorf("GiftService.GetAll scan: %w", err)
		}
		gifts = append(gifts, g)
	}
	return gifts, rows.Err()
}

// GetByID returns the gift with the given id, or pgx.ErrNoRows if not found.
func (s *GiftService) GetByID(id int) (*models.Gift, error) {
	var g models.Gift
	err := s.db.QueryRow(context.Background(), `
		SELECT id, img, name, link, progress, total, purchased, description
		FROM gifts WHERE id = $1`, id,
	).Scan(
		&g.ID, &g.Img, &g.Name, &g.Link,
		&g.Progress, &g.Total, &g.Purchased, &g.Description,
	)
	if err != nil {
		return nil, fmt.Errorf("GiftService.GetByID: %w", err)
	}
	return &g, nil
}

// Create inserts a new gift and returns the record with its generated id.
func (s *GiftService) Create(g *models.Gift) (*models.Gift, error) {
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO gifts (img, name, link, progress, total, purchased, description)
		VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`,
		g.Img, g.Name, g.Link, g.Progress, g.Total, g.Purchased, g.Description,
	).Scan(&g.ID)
	if err != nil {
		return nil, fmt.Errorf("GiftService.Create: %w", err)
	}
	return g, nil
}

// Update overwrites all mutable fields for the gift with the given id.
func (s *GiftService) Update(id int, g *models.Gift) (*models.Gift, error) {
	tag, err := s.db.Exec(context.Background(), `
		UPDATE gifts SET
			img         = $1,
			name        = $2,
			link        = $3,
			progress    = $4,
			total       = $5,
			purchased   = $6,
			description = $7
		WHERE id = $8`,
		g.Img, g.Name, g.Link, g.Progress, g.Total, g.Purchased, g.Description, id,
	)
	if err != nil {
		return nil, fmt.Errorf("GiftService.Update: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}
	g.ID = id
	return g, nil
}

// Delete removes the gift with the given id.
func (s *GiftService) Delete(id int) error {
	tag, err := s.db.Exec(context.Background(), `DELETE FROM gifts WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("GiftService.Delete: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
