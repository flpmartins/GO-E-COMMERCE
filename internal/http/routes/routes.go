package routes

import (
	"tmp-api/internal/http/handlers"

	"github.com/go-chi/chi/v5"
)

func NewRouter(userHandler *handlers.UserHandler) *chi.Mux {
	routes := chi.NewRouter()

	routes.Route("/users", func(r chi.Router) {
		routes.Post("/", userHandler.CreateUser)
		routes.Get("/", userHandler.GetAllUsers)
		routes.Get("/{id}", userHandler.GetUser)
		routes.Put("/{id}", userHandler.UpdateUser)
		routes.Delete("/{id}", userHandler.DeleteUser)
	})

	return routes
}
