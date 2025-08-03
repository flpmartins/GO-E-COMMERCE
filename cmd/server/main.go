package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tmp-api/internal/configs"
	"tmp-api/internal/http/handlers"
	"tmp-api/internal/http/routes"
	"tmp-api/internal/repository"
	"tmp-api/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Adicionando métodos...")

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env não encontrado, usando valores padrão.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := configs.NewPostgresConnection()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco: ", err)
	}

	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	permissionRepo := repository.NewPermissionRepository(db)
	permissionService := service.NewPermissionService(permissionRepo)
	permissionHandler := handlers.NewPermissionHandler(permissionService)

	router := routes.NewRouter(userHandler, permissionHandler)

	fmt.Println("Servidor rodando na porta", port, "🚀⭐")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Erro ao iniciar servidor: ", err)
	}
}
