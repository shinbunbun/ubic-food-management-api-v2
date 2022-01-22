package user

type User struct {
	UserID       string        `json:"userId"`
	Name         string        `json:"name"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	ID   string `json:"id"`
	Date string `json:"date"`
	Food Food   `json:"food"`
}

type Food struct {
	ID       string `json:"id"`
	ImageUrl string `json:"imageUrl"`
	Maker    string `json:"maker"`
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
}
