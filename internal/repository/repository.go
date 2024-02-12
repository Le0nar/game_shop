package repository

import (
	"fmt"

	"github.com/Le0nar/game_shop/internal/user"
	"github.com/jmoiron/sqlx"
)

const accountTable = "account"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

// Create a user with unic id and 0 gold
func (r *Repository) CreateUser(nickname string) error {
	query := fmt.Sprintf("INSERT INTO %s (name) values ($1)", accountTable)

	_, err := r.db.Exec(query, nickname)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUser(id int) (*user.User, error) {
	var user user.User

	query := fmt.Sprintf("SELECT * FROM %s where id = %d", accountTable, id)
	err := r.db.Get(&user, query)

	return &user, err
}

// Add gold to account
func (r *Repository) AddGold(nickname string, quantity int) error {
	return nil
}

// Buy item, if user have enough gold
func (r *Repository) BuyItem(itemId string, quantity int) error {
	return nil
}
