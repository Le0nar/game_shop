package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Le0nar/game_shop/internal/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type service interface {
	CreateUser(nickname string) error
	AddGold(nickname string, quantity int) error
	BuyItem(itemId string, quantity int) error
	RefundItem(itemId string, quantity int) error
	GetGold(nickname string) (int, error)
}

type Handler struct {
	service service
}

func NewHandler(s service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user user.UserDTO

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

// func (h *Handler) AddGold(nickname string, quantity int) error {
// 	return h.service.AddGold(nickname, quantity)
// }

// func (h *Handler) BuyItem(itemId string, quantity int) error {
// 	return h.service.BuyItem(itemId, quantity)
// }

// func (h *Handler) RefundItem(itemId string, quantity int) error {
// 	// user have 15 minutes for refound item
// 	return h.service.RefundItem(itemId, quantity)
// }

func (h *Handler) GetGold(w http.ResponseWriter, r *http.Request) {
	nickname := r.URL.Query().Get("nickname")
	if nickname == "" {
		http.Error(w, "nickname query param is missing", http.StatusBadRequest)
		return
	}

	quantity, err := h.service.GetGold(nickname)
	if err != nil {
		// TODO: change return value
		render.JSON(w, r, err.Error())
		return
	}

	// TODO: return struct with {qunatity: value}

	render.JSON(w, r, quantity)
}

// Initialization of router
func (h *Handler) InitRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/user", h.CreateUser)
	router.Get("/gold", h.GetGold)

	return router
}
