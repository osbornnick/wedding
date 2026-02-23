// GuestController handles HTTP requests for the /guests resource.
package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	"wedding/services/models"
	"wedding/services/services"
)

// GuestController holds a reference to GuestService and exposes handler methods.
type GuestController struct {
	svc *services.GuestService
}

// NewGuestController constructs a GuestController.
func NewGuestController(svc *services.GuestService) *GuestController {
	return &GuestController{svc: svc}
}

// GetAll handles GET /guests — returns all guests as a JSON array.
func (c *GuestController) GetAll(w http.ResponseWriter, r *http.Request) {
	guests, err := c.svc.GetAll()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Return an empty array rather than null when there are no guests.
	if guests == nil {
		guests = []models.Guest{}
	}
	writeJSON(w, http.StatusOK, guests)
}

// GetByID handles GET /guests/{id} — returns a single guest by id.
func (c *GuestController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be an integer")
		return
	}
	guest, err := c.svc.GetByID(id)
	if err == pgx.ErrNoRows {
		writeError(w, http.StatusNotFound, "guest not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, guest)
}

// Create handles POST /guests — creates a new guest from the JSON body.
func (c *GuestController) Create(w http.ResponseWriter, r *http.Request) {
	var g models.Guest
	if !decodeJSON(w, r, &g) {
		return
	}
	created, err := c.svc.Create(&g)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// Update handles PUT /guests/{id} — replaces an existing guest with the JSON body.
func (c *GuestController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be an integer")
		return
	}
	var g models.Guest
	if !decodeJSON(w, r, &g) {
		return
	}
	updated, err := c.svc.Update(id, &g)
	if err == pgx.ErrNoRows {
		writeError(w, http.StatusNotFound, "guest not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

// Delete handles DELETE /guests/{id} — removes a guest by id.
func (c *GuestController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be an integer")
		return
	}
	if err := c.svc.Delete(id); err == pgx.ErrNoRows {
		writeError(w, http.StatusNotFound, "guest not found")
		return
	} else if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *GuestController) FuzzySearchByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		writeError(w, http.StatusBadRequest, "name query parameter is required")
		return
	}
	results, err := c.svc.FuzzySearchByName(name)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, results)
}
