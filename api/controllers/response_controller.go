// ResponseController handles HTTP requests for the /responses resource.
package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	"wedding/services/models"
	"wedding/services/services"
)

// ResponseController holds a reference to ResponseService and exposes handler methods.
type ResponseController struct {
	svc *services.ResponseService
}

// NewResponseController constructs a ResponseController.
func NewResponseController(svc *services.ResponseService) *ResponseController {
	return &ResponseController{svc: svc}
}

// GetAll handles GET /responses — returns all RSVP responses as a JSON array.
func (c *ResponseController) GetAll(w http.ResponseWriter, r *http.Request) {
	responses, err := c.svc.GetAll()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if responses == nil {
		responses = []models.Response{}
	}
	writeJSON(w, http.StatusOK, responses)
}

// GetByID handles GET /responses/{id} — returns a single RSVP response by id.
func (c *ResponseController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be an integer")
		return
	}
	resp, err := c.svc.GetByID(id)
	if err == pgx.ErrNoRows {
		writeError(w, http.StatusNotFound, "response not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

// Create handles POST /responses — creates a new RSVP response from the JSON body.
func (c *ResponseController) Create(w http.ResponseWriter, r *http.Request) {
	var resp models.Response
	if !decodeJSON(w, r, &resp) {
		return
	}
	created, err := c.svc.Create(&resp)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// Update handles PUT /responses/{id} — replaces an existing RSVP response with the JSON body.
func (c *ResponseController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be an integer")
		return
	}
	var resp models.Response
	if !decodeJSON(w, r, &resp) {
		return
	}
	updated, err := c.svc.Update(id, &resp)
	if err == pgx.ErrNoRows {
		writeError(w, http.StatusNotFound, "response not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

// Delete handles DELETE /responses/{id} — removes an RSVP response by id.
func (c *ResponseController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be an integer")
		return
	}
	if err := c.svc.Delete(id); err == pgx.ErrNoRows {
		writeError(w, http.StatusNotFound, "response not found")
		return
	} else if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
