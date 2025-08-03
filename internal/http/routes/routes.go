package routes

import (
	"tmp-api/internal/http/handlers"

	"github.com/go-chi/chi/v5"
)

func NewRouter(userHandler *handlers.UserHandler, permissionHandler *handlers.PermissionHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Get("/", userHandler.GetAllUsers)
		r.Get("/{id}", userHandler.GetUser)
		r.Put("/{id}", userHandler.UpdateUser)
		r.Delete("/{id}", userHandler.DeleteUser)
	})

	router.Route("/permissions", func(r chi.Router) {
		r.Post("/", permissionHandler.CreatePermission)
		r.Get("/", permissionHandler.GetAllPermissions)
		r.Get("/{id}", permissionHandler.GetPermission)
		r.Put("/{id}", permissionHandler.UpdatePermission)
		r.Delete("/{id}", permissionHandler.DeletePermission)
	})

	return router
}
