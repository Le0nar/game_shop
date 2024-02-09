package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

// func (h *Handler) CreateUser(nickname string) error {
// 	return h.service.CreateUser(nickname)
// }

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

	w.Write([]byte(nickname))
}

// Initialization of router
func (h *Handler) InitRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/get-gold", h.GetGold)

	return router
}
