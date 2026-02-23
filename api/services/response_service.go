// ResponseService provides CRUD operations for the responses table.
package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"wedding/services/models"
)

// ResponseService wraps the connection pool and exposes RSVP response operations.
type ResponseService struct {
	db *pgxpool.Pool
}

// NewResponseService constructs a ResponseService with the given connection pool.
func NewResponseService(db *pgxpool.Pool) *ResponseService {
	return &ResponseService{db: db}
}

// GetAll returns every RSVP response from the database.
func (s *ResponseService) GetAll() ([]models.Response, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT id,
		       attending_wedding, attending_wedding_count,
		       attending_friday,  attending_friday_count,
		       attending_brunch,  attending_brunch_count,
		       dietary_restrictions
		FROM responses ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("ResponseService.GetAll query: %w", err)
	}
	defer rows.Close()

	var responses []models.Response
	for rows.Next() {
		var r models.Response
		if err := rows.Scan(
			&r.ID,
			&r.AttendingWedding, &r.AttendingWeddingCount,
			&r.AttendingFriday, &r.AttendingFridayCount,
			&r.AttendingBrunch, &r.AttendingBrunchCount,
			&r.DietaryRestrictions,
		); err != nil {
			return nil, fmt.Errorf("ResponseService.GetAll scan: %w", err)
		}
		responses = append(responses, r)
	}
	return responses, rows.Err()
}

// GetByID returns the response with the given id, or pgx.ErrNoRows if not found.
func (s *ResponseService) GetByID(id int) (*models.Response, error) {
	var r models.Response
	err := s.db.QueryRow(context.Background(), `
		SELECT id,
		       attending_wedding, attending_wedding_count,
		       attending_friday,  attending_friday_count,
		       attending_brunch,  attending_brunch_count,
		       dietary_restrictions
		FROM responses WHERE id = $1`, id,
	).Scan(
		&r.ID,
		&r.AttendingWedding, &r.AttendingWeddingCount,
		&r.AttendingFriday, &r.AttendingFridayCount,
		&r.AttendingBrunch, &r.AttendingBrunchCount,
		&r.DietaryRestrictions,
	)
	if err != nil {
		return nil, fmt.Errorf("ResponseService.GetByID: %w", err)
	}
	return &r, nil
}

// Create inserts a new RSVP response and returns the record with its generated id.
func (s *ResponseService) Create(r *models.Response) (*models.Response, error) {
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO responses (
			attending_wedding, attending_wedding_count,
			attending_friday,  attending_friday_count,
			attending_brunch,  attending_brunch_count,
			dietary_restrictions
		) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`,
		r.AttendingWedding, r.AttendingWeddingCount,
		r.AttendingFriday, r.AttendingFridayCount,
		r.AttendingBrunch, r.AttendingBrunchCount,
		r.DietaryRestrictions,
	).Scan(&r.ID)
	if err != nil {
		return nil, fmt.Errorf("ResponseService.Create: %w", err)
	}
	return r, nil
}

// Update overwrites all mutable fields for the response with the given id.
func (s *ResponseService) Update(id int, r *models.Response) (*models.Response, error) {
	tag, err := s.db.Exec(context.Background(), `
		UPDATE responses SET
			attending_wedding       = $1,
			attending_wedding_count = $2,
			attending_friday        = $3,
			attending_friday_count  = $4,
			attending_brunch        = $5,
			attending_brunch_count  = $6,
			dietary_restrictions    = $7
		WHERE id = $8`,
		r.AttendingWedding, r.AttendingWeddingCount,
		r.AttendingFriday, r.AttendingFridayCount,
		r.AttendingBrunch, r.AttendingBrunchCount,
		r.DietaryRestrictions, id,
	)
	if err != nil {
		return nil, fmt.Errorf("ResponseService.Update: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}
	r.ID = id
	return r, nil
}

// Delete removes the RSVP response with the given id.
func (s *ResponseService) Delete(id int) error {
	tag, err := s.db.Exec(context.Background(), `DELETE FROM responses WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("ResponseService.Delete: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
