// PurchaseService provides CRUD operations for the purchases table.
package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"wedding/services/models"
)

// PurchaseService wraps the connection pool and exposes purchase operations.
type PurchaseService struct {
	db *pgxpool.Pool
}

// NewPurchaseService constructs a PurchaseService with the given connection pool.
func NewPurchaseService(db *pgxpool.Pool) *PurchaseService {
	return &PurchaseService{db: db}
}

// GetAll returns every purchase record from the database.
func (s *PurchaseService) GetAll() ([]models.Purchase, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT id, gift_id, guest_id, amount, created_at, updated_at
		FROM purchases ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("PurchaseService.GetAll query: %w", err)
	}
	defer rows.Close()

	var purchases []models.Purchase
	for rows.Next() {
		var p models.Purchase
		if err := rows.Scan(
			&p.ID, &p.GiftID, &p.GuestID, &p.Amount, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("PurchaseService.GetAll scan: %w", err)
		}
		purchases = append(purchases, p)
	}
	return purchases, rows.Err()
}

// GetByID returns the purchase with the given id, or pgx.ErrNoRows if not found.
func (s *PurchaseService) GetByID(id int) (*models.Purchase, error) {
	var p models.Purchase
	err := s.db.QueryRow(context.Background(), `
		SELECT id, gift_id, guest_id, amount, created_at, updated_at
		FROM purchases WHERE id = $1`, id,
	).Scan(
		&p.ID, &p.GiftID, &p.GuestID, &p.Amount, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("PurchaseService.GetByID: %w", err)
	}
	return &p, nil
}

// Create inserts a new purchase and returns the record with its generated id and timestamps.
func (s *PurchaseService) Create(p *models.Purchase) (*models.Purchase, error) {
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO purchases (gift_id, guest_id, amount)
		VALUES ($1,$2,$3) RETURNING id, created_at, updated_at`,
		p.GiftID, p.GuestID, p.Amount,
	).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("PurchaseService.Create: %w", err)
	}
	return p, nil
}

// Update overwrites all mutable fields for the purchase with the given id.
func (s *PurchaseService) Update(id int, p *models.Purchase) (*models.Purchase, error) {
	err := s.db.QueryRow(context.Background(), `
		UPDATE purchases SET
			gift_id    = $1,
			guest_id   = $2,
			amount     = $3,
			updated_at = NOW()
		WHERE id = $4
		RETURNING updated_at`,
		p.GiftID, p.GuestID, p.Amount, id,
	).Scan(&p.UpdatedAt)
	if err == pgx.ErrNoRows {
		return nil, pgx.ErrNoRows
	}
	if err != nil {
		return nil, fmt.Errorf("PurchaseService.Update: %w", err)
	}
	p.ID = id
	return p, nil
}

// Delete removes the purchase with the given id.
func (s *PurchaseService) Delete(id int) error {
	tag, err := s.db.Exec(context.Background(), `DELETE FROM purchases WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("PurchaseService.Delete: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
