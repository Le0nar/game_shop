package user

type User struct {
	Id       int    `json:"id" db:"id"`
	Nickname string `json:"nickname" db:"name"`
	Gold     int    `json:"gold" db:"gold"`
}

type CreateUserDTO struct {
	Nickname string `json:"nickname"`
}

type AddGoldDTO struct {
	Id   int `json:"id" db:"id"`
	Gold int `json:"gold" db:"gold"`
}
