package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Le0nar/game_shop/internal/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type service interface {
	CreateUser(nickname string) error
	GetUser(id int) (*user.User, error)
	AddGold(nickname string, quantity int) error
	BuyItem(itemId string, quantity int) error
	RefundItem(itemId string, quantity int) error
}

type Handler struct {
	service service
}

func NewHandler(s service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user user.CreateUserDTO

	// TODO: add check for empty json
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateUser(user.Nickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, "")
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "invalid id param", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, *user)
}

// Initialization of router
func (h *Handler) InitRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/user", h.CreateUser)
	router.Get("/user", h.GetUser)

	return router
}
