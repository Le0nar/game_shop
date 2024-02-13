package repository

import (
	"fmt"

	"github.com/Le0nar/game_shop/internal/inventory"
	"github.com/Le0nar/game_shop/internal/item"
	"github.com/Le0nar/game_shop/internal/user"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const accountTable = "account"
const itemTable = "item"
const inventoryTable = "inventory"

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
func (r *Repository) AddGold(id int, quantity int) error {
	query := fmt.Sprintf("UPDATE %s SET gold = gold + %d WHERE id = %d", accountTable, quantity, id)

	_, err := r.db.Exec(query)

	return err
}

func (r *Repository) GetItem(id int) (*item.Item, error) {
	var item item.Item

	query := fmt.Sprintf("SELECT * FROM %s where id = %d", itemTable, id)
	err := r.db.Get(&item, query)

	return &item, err
}

// TODO: mb add tx.Rollback()
// Buy item, if user have enough gold
func (r *Repository) BuyItem(user *user.User, item *item.Item) error {
	// create transactions:
	tx, err := r.db.Begin()

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	if err != nil {
		return err
	}

	// 1) decrase user's gold
	// TODO: mb add  "WHERE id = %d AND gold >= item.Price"
	query := fmt.Sprintf("UPDATE %s SET gold = gold - %d WHERE id = %d", accountTable, item.Price, user.Id)

	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	// 2) add itemId to user's inventory
	var inventory inventory.Inventory

	query = fmt.Sprintf("SELECT * FROM %s where accountId = %d", inventoryTable, user.Id)
	err = r.db.Get(&user, query)

	// if user's inventory doesnt exist, create user's inventory
	if err != nil {
		query := fmt.Sprintf("INSERT INTO %s values ($1, $2)", inventoryTable)
		itemIdList := []int{item.Id}

		_, err := r.db.Exec(query, user.Id, pq.Array(itemIdList))
		if err != nil {
			return err
		}
		// if user has inventory
	} else {
		inventory.ItemIdList = append(inventory.ItemIdList, item.Id)

		fmt.Printf("inventory.ItemIdList: %v\n", inventory.ItemIdList)

		query := "UPDATE $1 SET itemIdList = ANY($2) WHERE id = $3"

		_, err = tx.Exec(query, inventoryTable, pq.Array(inventory.ItemIdList), user.Id)
		if err != nil {
			return err
		}

	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
