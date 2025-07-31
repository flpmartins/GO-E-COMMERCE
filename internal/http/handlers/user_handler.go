package handlers

import (
	"fmt"
	"net/http"
	"tmp-api/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Obtendo usuário!")
	return
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Obtendo usuário!")
	return
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Listando todos os usuários!")
	return
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Atualizando usuário!")
	return
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Removendo usuário!")
	return
}
