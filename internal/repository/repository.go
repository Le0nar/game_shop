package repository

type Repository struct {
	// db *sqlx.DB
}

func NewRepository() *Repository {
	// func NewRepository(db *sqlx.DB) *Repository {
	// return &Repository{db: db}
	return &Repository{}
}

// Create a user with unic nickname
func (r *Repository) CreateUser(nickname string) error {
	return nil
}

// Add gold to account
func (r *Repository) AddGold(nickname string, quantity int) error {
	return nil
}

// Buy item, if user have enough gold
func (r *Repository) BuyItem(itemId string, quantity int) error {
	return nil
}

// Get current gold on accound
func (r *Repository) GetGold(nickname string) (int, error) {
	return 0, nil
}
