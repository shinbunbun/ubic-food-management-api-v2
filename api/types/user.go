package types

type User struct {
	UserID       string        `json:"userId"`
	Name         string        `json:"name"`
	Transactions []Transaction `json:"transactions"`
}
