package service

import "github.com/Le0nar/game_shop/internal/user"

type repository interface {
	CreateUser(nickname string) error
	GetUser(id int) (*user.User, error)
	AddGold(nickname string, quantity int) error
	BuyItem(itemId string, quantity int) error
}

type Serivce struct {
	repository repository
}

func NewService(r repository) *Serivce {
	return &Serivce{repository: r}
}

func (s *Serivce) CreateUser(nickname string) error {
	return s.repository.CreateUser(nickname)
}

func (s *Serivce) GetUser(id int) (*user.User, error) {
	return s.repository.GetUser(id)
}

func (s *Serivce) AddGold(nickname string, quantity int) error {
	return s.repository.AddGold(nickname, quantity)
}

func (s *Serivce) BuyItem(itemId string, quantity int) error {
	return s.repository.BuyItem(itemId, quantity)
}

func (s *Serivce) RefundItem(itemId string, quantity int) error {
	// user have 15 minutes for refound item
	return s.repository.AddGold(itemId, quantity)
}
