// InvitationController handles HTTP requests for the /invitations resource.
package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	"wedding/services/models"
	"wedding/services/services"
)

// InvitationController holds a reference to InvitationService and exposes handler methods.
type InvitationController struct {
	svc *services.InvitationService
}

// NewInvitationController constructs an InvitationController.
func NewInvitationController(svc *services.InvitationService) *InvitationController {
	return &InvitationController{svc: svc}
}

// GetAll handles GET /invitations — returns all invitations as a JSON array.
func (c *InvitationController) GetAll(w http.ResponseWriter, r *http.Request) {
	invitations, err := c.svc.GetAll()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if invitations == nil {
		invitations = []models.Invitation{}
	}
	writeJSON(w, http.StatusOK, invitations)
}

// GetByID handles GET /invitations/{id} — returns a single invitation by id.
func (c *InvitationController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be an integer")
		return
	}
	inv, err := c.svc.GetByID(id)
	if err == pgx.ErrNoRows {
		writeError(w, http.StatusNotFound, "invitation not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, inv)
}

// Create handles POST /invitations — creates a new invitation from the JSON body.
func (c *InvitationController) Create(w http.ResponseWriter, r *http.Request) {
	var inv models.Invitation
	if !decodeJSON(w, r, &inv) {
		return
	}
	created, err := c.svc.Create(&inv)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// Update handles PUT /invitations/{id} — replaces an existing invitation with the JSON body.
func (c *InvitationController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be an integer")
		return
	}
	var inv models.Invitation
	if !decodeJSON(w, r, &inv) {
		return
	}
	updated, err := c.svc.Update(id, &inv)
	if err == pgx.ErrNoRows {
		writeError(w, http.StatusNotFound, "invitation not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

// Delete handles DELETE /invitations/{id} — removes an invitation by id.
func (c *InvitationController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be an integer")
		return
	}
	if err := c.svc.Delete(id); err == pgx.ErrNoRows {
		writeError(w, http.StatusNotFound, "invitation not found")
		return
	} else if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
