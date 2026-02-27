package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) *UserService {
	return &UserService{db: db}
}

func (s *UserService) Login(username, password string) (bool, error) {
	var storedPassword string
	err := s.db.QueryRow(
		context.Background(),
		`SELECT password_hash FROM users WHERE username = $1`, username,
	).Scan(&storedPassword)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (s *UserService) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
