package handlers

import (
	"html/template"
	"net/http"
	"time"

	"github.com/ryantrue/contractkeeper/internal/models"
	"github.com/ryantrue/contractkeeper/internal/services"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	service services.UserService
	logger  *logrus.Logger
}

func NewUserHandler(service services.UserService, logger *logrus.Logger) *UserHandler {
	return &UserHandler{service: service, logger: logger}
}

func (h *UserHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		h.logger.Errorf("Could not load template: %v", err)
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (h *UserHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		h.logger.Errorf("Could not load template: %v", err)
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := h.service.Authenticate(username, password)
	if err != nil || user == nil {
		h.logger.Warn("Invalid login attempt")
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "example_token", // generate a real token here
		Expires: time.Now().Add(24 * time.Hour),
	})

	h.logger.Infof("User %s logged in", username)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user := models.User{
		Username: username,
		Password: password,
	}

	err := h.service.CreateUser(&user)
	if err != nil {
		h.logger.Errorf("Could not register user: %v", err)
		http.Error(w, "Could not register user", http.StatusInternalServerError)
		return
	}

	h.logger.Infof("User %s registered", username)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "session_token",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	h.logger.Info("User logged out")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *UserHandler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}
