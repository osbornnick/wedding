// GiftController handles HTTP requests for the /gifts resource.
package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	"wedding/services/models"
	"wedding/services/services"
)

// GiftController holds a reference to GiftService and exposes handler methods.
type GiftController struct {
	svc *services.GiftService
}

// NewGiftController constructs a GiftController.
func NewGiftController(svc *services.GiftService) *GiftController {
	return &GiftController{svc: svc}
}

// GetAll handles GET /gifts — returns all gift registry items as a JSON array.
func (c *GiftController) GetAll(w http.ResponseWriter, r *http.Request) {
	gifts, err := c.svc.GetAll()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if gifts == nil {
		gifts = []models.Gift{}
	}
	writeJSON(w, http.StatusOK, gifts)
}

// GetByID handles GET /gifts/{id} — returns a single gift by id.
func (c *GiftController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be an integer")
		return
	}
	gift, err := c.svc.GetByID(id)
	if err == pgx.ErrNoRows {
		writeError(w, http.StatusNotFound, "gift not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, gift)
}

// Create handles POST /gifts — creates a new gift registry item from the JSON body.
func (c *GiftController) Create(w http.ResponseWriter, r *http.Request) {
	var g models.Gift
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

// Update handles PUT /gifts/{id} — replaces an existing gift with the JSON body.
func (c *GiftController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be an integer")
		return
	}
	var g models.Gift
	if !decodeJSON(w, r, &g) {
		return
	}
	updated, err := c.svc.Update(id, &g)
	if err == pgx.ErrNoRows {
		writeError(w, http.StatusNotFound, "gift not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

// Delete handles DELETE /gifts/{id} — removes a gift by id.
func (c *GiftController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be an integer")
		return
	}
	if err := c.svc.Delete(id); err == pgx.ErrNoRows {
		writeError(w, http.StatusNotFound, "gift not found")
		return
	} else if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
