package types

type Transaction struct {
	ID   string `json:"id"`
	Date string `json:"date"`
	Food Food   `json:"food"`
}
