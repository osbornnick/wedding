// Package router wires all controllers to their URL paths using the chi mux.
package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"wedding/services/controllers"
)

// New creates and returns a fully configured chi.Mux.
// Each entity gets a sub-router at /api/<resource> with standard CRUD routes:
//
//	GET    /api/<resource>      → list all
//	POST   /api/<resource>      → create
//	GET    /api/<resource>/{id} → get one
//	PUT    /api/<resource>/{id} → replace one
//	DELETE /api/<resource>/{id} → delete one
func New(
	guestCtrl *controllers.GuestController,
	invitationCtrl *controllers.InvitationController,
	responseCtrl *controllers.ResponseController,
	giftCtrl *controllers.GiftController,
	userCtrl *controllers.UserController,
) http.Handler {
	r := chi.NewRouter()

	// ── Global middleware ────────────────────────────────────────────────────────
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:*", "https://localhost:*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))
	r.Use(middleware.Logger)    // log every request with method, path, status, latency
	r.Use(middleware.Recoverer) // recover from panics and return 500 instead of crashing

	// ── Health check ─────────────────────────────────────────────────────────────
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// ── API routes ───────────────────────────────────────────────────────────────
	r.Route("/api", func(r chi.Router) {

		// Guests
		r.Route("/guests", func(r chi.Router) {
			r.Get("/search", guestCtrl.FuzzySearchByName)
			r.Get("/", guestCtrl.GetAll)
			r.Post("/", guestCtrl.Create)
			r.Get("/{id}", guestCtrl.GetByID)
			r.Put("/{id}", guestCtrl.Update)
			r.Delete("/{id}", guestCtrl.Delete)
		})

		// Invitations
		r.Route("/invitations", func(r chi.Router) {
			r.Get("/", invitationCtrl.GetAll)
			r.Post("/", invitationCtrl.Create)
			r.Get("/{id}", invitationCtrl.GetByID)
			r.Put("/{id}", invitationCtrl.Update)
			r.Delete("/{id}", invitationCtrl.Delete)
		})

		// Responses (RSVP)
		r.Route("/responses", func(r chi.Router) {
			r.Get("/", responseCtrl.GetAll)
			r.Post("/", responseCtrl.Create)
			r.Get("/{id}", responseCtrl.GetByID)
			r.Put("/{id}", responseCtrl.Update)
			r.Delete("/{id}", responseCtrl.Delete)
		})

		// Gifts (registry)
		r.Route("/gifts", func(r chi.Router) {
			r.Get("/", giftCtrl.GetAll)
			r.Post("/", giftCtrl.Create)
			r.Get("/{id}", giftCtrl.GetByID)
			r.Put("/{id}", giftCtrl.Update)
			r.Delete("/{id}", giftCtrl.Delete)
		})

		r.Route("/users", func(r chi.Router) {
			r.Post("/login", userCtrl.Login)
			r.Post("/hash", userCtrl.HashPassword)
		})
	})

	return r
}
