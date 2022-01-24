package types

type Food struct {
	ID       string `json:"id"`
	ImageUrl string `json:"imageUrl"`
	Maker    string `json:"maker"`
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
}
