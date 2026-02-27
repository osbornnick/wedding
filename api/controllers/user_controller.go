package controllers

import (
	"fmt"
	"net/http"
	"wedding/services/services"
)

type UserController struct {
	svc *services.UserService
}

func NewUserController(svc *services.UserService) *UserController {
	return &UserController{svc: svc}
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if !decodeJSON(w, r, &creds) {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
	success, err := c.svc.Login(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("login failed: %v", err), http.StatusInternalServerError)
		return
	}
	if !success {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *UserController) HashPassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Password string `json:"password"`
	}
	if !decodeJSON(w, r, &req) {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
	hashed, err := c.svc.HashPassword(req.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to hash password: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"hashed_password":"%s"}`, hashed)))
}