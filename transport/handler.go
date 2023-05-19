package transport

import (
	"net/http"
	"newsletter-platform/transport/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

type Handler struct {
	Port     int
	Mux      *chi.Mux
	Services model.ServiceCollection
}

// Initializer function for HTTP handler.
func Initialize(port int, services model.ServiceCollection) *Handler {
	handler := &Handler{
		Port:     port,
		Mux:      chi.NewRouter(),
		Services: services,
	}
	handler.Mux.Use(middleware.Logger)

	// Protected routes
	handler.Mux.Group(func(r chi.Router) {
		//r.Use(jwtauth.Verifier(handler.AuthService.GetJwtAuth()))
		r.Use(jwtauth.Authenticator)
	})

	// Public routes
	handler.Mux.Get("/health", health)
	handler.Mux.Post("/register", handler.RegisterUser)

	return handler
}

// Controller endpoint function for handling requests on /health.
func health(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNoContent)
}
