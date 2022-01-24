package types

type Transaction struct {
	ID   string `json:"id"`
	Date int    `json:"date"`
	Food Food   `json:"food"`
}
