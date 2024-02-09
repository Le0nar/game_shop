package service

type repository interface {
	CreateUser(nickname string) error
	AddGold(nickname string, quantity int) error
	BuyItem(itemId string, quantity int) error
	GetGold(nickname string) (int, error)
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

func (s *Serivce) GetGold(nickname string) (int, error) {
	return s.repository.GetGold(nickname)
}
