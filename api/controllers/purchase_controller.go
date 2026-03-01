// PurchaseController handles HTTP requests for the /purchases resource.
package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	"wedding/services/models"
	"wedding/services/services"
)

// PurchaseController holds a reference to PurchaseService and exposes handler methods.
type PurchaseController struct {
	svc *services.PurchaseService
}

// NewPurchaseController constructs a PurchaseController.
func NewPurchaseController(svc *services.PurchaseService) *PurchaseController {
	return &PurchaseController{svc: svc}
}

// GetAll handles GET /purchases — returns all purchase records as a JSON array.
func (c *PurchaseController) GetAll(w http.ResponseWriter, r *http.Request) {
	purchases, err := c.svc.GetAll()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if purchases == nil {
		purchases = []models.Purchase{}
	}
	writeJSON(w, http.StatusOK, purchases)
}

// GetByID handles GET /purchases/{id} — returns a single purchase by id.
func (c *PurchaseController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be an integer")
		return
	}
	purchase, err := c.svc.GetByID(id)
	if err == pgx.ErrNoRows {
		writeError(w, http.StatusNotFound, "purchase not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, purchase)
}

// Create handles POST /purchases — creates a new purchase from the JSON body.
func (c *PurchaseController) Create(w http.ResponseWriter, r *http.Request) {
	var p models.Purchase
	if !decodeJSON(w, r, &p) {
		return
	}
	created, err := c.svc.Create(&p)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// Update handles PUT /purchases/{id} — replaces an existing purchase with the JSON body.
func (c *PurchaseController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be an integer")
		return
	}
	var p models.Purchase
	if !decodeJSON(w, r, &p) {
		return
	}
	updated, err := c.svc.Update(id, &p)
	if err == pgx.ErrNoRows {
		writeError(w, http.StatusNotFound, "purchase not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

// Delete handles DELETE /purchases/{id} — removes a purchase by id.
func (c *PurchaseController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be an integer")
		return
	}
	if err := c.svc.Delete(id); err == pgx.ErrNoRows {
		writeError(w, http.StatusNotFound, "purchase not found")
		return
	} else if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
