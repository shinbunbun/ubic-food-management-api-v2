package types

import (
	"hello-world/dynamodb"
)

type Food struct {
	ID       string `json:"id"`
	ImageUrl string `json:"imageUrl"`
	Maker    string `json:"maker"`
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
}

func (f *Food) Get() error {

	items, err := dynamodb.GetByID(f.ID)
	if err != nil {
		return err
	}

	for _, v := range items {
		if v.DataType == "food-image" {
			f.ImageUrl = v.Data
		}
		if v.DataType == "food-maker" {
			f.Maker = v.Data
		}
		if v.DataType == "food-name" {
			f.Name = v.Data
		}
		if v.DataType == "food-stock" {
			f.Stock = v.IntData
		}
	}
	return nil
}
