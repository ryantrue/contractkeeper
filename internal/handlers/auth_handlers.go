package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/ryantrue/contractkeeper/internal/services"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	service services.UserService
	logger  *logrus.Logger
}

func NewAuthHandler(service services.UserService, logger *logrus.Logger) *AuthHandler {
	return &AuthHandler{service: service, logger: logger}
}

func (h *AuthHandler) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "login.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		h.logger.Errorf("Could not load template: %v", err)
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (h *AuthHandler) RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "register.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		h.logger.Errorf("Could not load template: %v", err)
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := h.service.Authenticate(username, password)
	if err != nil {
		h.logger.Errorf("Error during authentication: %v", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	session, _ := store.Get(r, "session")
	session.Values["user_id"] = user.ID
	session.Save(r, w)

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	err := h.service.Register(username, password)
	if err != nil {
		h.logger.Errorf("Error during registration: %v", err)
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	delete(session.Values, "user_id")
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *AuthHandler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		userID, ok := session.Values["user_id"].(uint)
		if !ok || userID == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	}
}
