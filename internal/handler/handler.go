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

	AddGold(id int, quantity int) error

	// TODO: add user id for methods
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
	var dto user.CreateUserDTO

	// TODO: add check for empty json
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateUser(dto.Nickname)
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

func (h *Handler) AddGold(w http.ResponseWriter, r *http.Request) {
	var dto user.AddGoldDTO

	// TODO: add check for empty json
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.AddGold(dto.Id, dto.Gold)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, "")
}

// Initialization of router
func (h *Handler) InitRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/user", h.CreateUser)
	router.Get("/user", h.GetUser)

	router.Patch("/gold", h.AddGold)

	return router
}
