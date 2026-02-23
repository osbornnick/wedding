// InvitationService provides CRUD operations for the invitations table.
package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"wedding/services/models"
)

// InvitationService wraps the connection pool and exposes invitation-related operations.
type InvitationService struct {
	db *pgxpool.Pool
}

// NewInvitationService constructs an InvitationService with the given connection pool.
func NewInvitationService(db *pgxpool.Pool) *InvitationService {
	return &InvitationService{db: db}
}

// GetAll returns every invitation record from the database.
func (s *InvitationService) GetAll() ([]models.Invitation, error) {
	rows, err := s.db.Query(
		context.Background(),
		`SELECT id, address, num_guests, created_at FROM invitations ORDER BY id`,
	)
	if err != nil {
		return nil, fmt.Errorf("InvitationService.GetAll query: %w", err)
	}
	defer rows.Close()

	var invitations []models.Invitation
	for rows.Next() {
		var inv models.Invitation
		if err := rows.Scan(&inv.ID, &inv.Address, &inv.NumGuests, &inv.CreatedAt); err != nil {
			return nil, fmt.Errorf("InvitationService.GetAll scan: %w", err)
		}
		invitations = append(invitations, inv)
	}
	return invitations, rows.Err()
}

// GetByID returns the invitation with the given id, or pgx.ErrNoRows if not found.
func (s *InvitationService) GetByID(id int) (*models.Invitation, error) {
	var inv models.Invitation
	err := s.db.QueryRow(
		context.Background(),
		`SELECT id, address, num_guests, created_at FROM invitations WHERE id = $1`, id,
	).Scan(&inv.ID, &inv.Address, &inv.NumGuests, &inv.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("InvitationService.GetByID: %w", err)
	}
	return &inv, nil
}

// Create inserts a new invitation and returns the record with its generated id.
func (s *InvitationService) Create(inv *models.Invitation) (*models.Invitation, error) {
	err := s.db.QueryRow(
		context.Background(),
		`INSERT INTO invitations (address, num_guests)
		 VALUES ($1, $2) RETURNING id`,
		inv.Address, inv.NumGuests,
	).Scan(&inv.ID)
	if err != nil {
		return nil, fmt.Errorf("InvitationService.Create: %w", err)
	}
	return inv, nil
}

// Update overwrites all mutable fields for the invitation with the given id.
func (s *InvitationService) Update(id int, inv *models.Invitation) (*models.Invitation, error) {
	tag, err := s.db.Exec(
		context.Background(),
		`UPDATE invitations SET address = $1, num_guests = $2 WHERE id = $3`,
		inv.Address, inv.NumGuests, id,
	)
	if err != nil {
		return nil, fmt.Errorf("InvitationService.Update: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}
	inv.ID = id
	return inv, nil
}

// Delete removes the invitation with the given id.
func (s *InvitationService) Delete(id int) error {
	tag, err := s.db.Exec(context.Background(), `DELETE FROM invitations WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("InvitationService.Delete: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
