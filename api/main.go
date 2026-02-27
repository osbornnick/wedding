// main.go is the entry point for the wedding backend service.
// It initialises the database connection, wires up all controllers,
// registers routes, and starts the HTTP server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"wedding/services/controllers"
	"wedding/services/db"
	"wedding/services/router"
	"wedding/services/services"
)

func main() {
	// ── Database ────────────────────────────────────────────────────────────────
	// Build the Postgres DSN from environment variables so that secrets are never
	// hard-coded.  Defaults are provided for local development with Docker Compose.
	// pgx accepts a standard Postgres connection URL: postgres://user:pass@host:port/db
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_NAME", "wedding"),
	)
	database, err := db.Connect(dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Close()

	// ── Services ────────────────────────────────────────────────────────────────
	// Services contain all business logic and data-access for each entity.
	guestSvc := services.NewGuestService(database)
	invitationSvc := services.NewInvitationService(database)
	responseSvc := services.NewResponseService(database)
	giftSvc := services.NewGiftService(database)
	userSvc := services.NewUserService(database)

	// ── Controllers ─────────────────────────────────────────────────────────────
	// Controllers parse HTTP requests, delegate to services, and write responses.
	guestCtrl := controllers.NewGuestController(guestSvc)
	invitationCtrl := controllers.NewInvitationController(invitationSvc)
	responseCtrl := controllers.NewResponseController(responseSvc)
	giftCtrl := controllers.NewGiftController(giftSvc)
	userCtrl := controllers.NewUserController(userSvc)

	// ── Router ──────────────────────────────────────────────────────────────────
	r := router.New(guestCtrl, invitationCtrl, responseCtrl, giftCtrl, userCtrl)

	addr := ":" + getEnv("PORT", "8080")
	log.Printf("wedding backend listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

// getEnv returns the value of the environment variable named by key,
// or fallback if the variable is not set.
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
