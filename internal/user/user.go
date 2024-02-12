package user

type User struct {
	Id       int    `json:"id" db:"id"`
	Nickname string `json:"nickname" db:"name"`
	Gold     int    `json:"gold" db:"gold"`
}

type UserDTO struct {
	Nickname string `json:"nickname"`
}
