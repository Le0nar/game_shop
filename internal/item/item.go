package item

type Item struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Price int    `json:"price" db:"price"`
}

type BuyItemDTO struct {
	UserId int `json:"userId"`
	ItemId int `json:"itemId"`
}
