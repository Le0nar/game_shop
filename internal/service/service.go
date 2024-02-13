package service

import (
	"errors"

	"github.com/Le0nar/game_shop/internal/item"
	"github.com/Le0nar/game_shop/internal/user"
)

type repository interface {
	CreateUser(nickname string) error
	GetUser(id int) (*user.User, error)
	AddGold(id int, quantity int) error
	GetItem(id int) (*item.Item, error)
	BuyItem(user *user.User, item *item.Item) error
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

func (s *Serivce) AddGold(id int, quantity int) error {
	return s.repository.AddGold(id, quantity)
}

func (s *Serivce) BuyItem(userId, itemId int) error {
	// 1) check if item exist
	item, err := s.repository.GetItem(itemId)
	if err != nil {
		return err
	}

	// 2) check if user exist
	user, err := s.repository.GetUser(userId)
	if err != nil {
		return err
	}

	// 2.1) check if user have enough money
	if user.Gold < item.Price {
		return errors.New("Not enough gold")
	}

	return s.repository.BuyItem(user, item)
}
