package types

import (
	"hello-world/dynamodb"
	"strconv"
)

type Transaction struct {
	ID   string `json:"id"`
	Date int    `json:"date"`
	Food Food   `json:"food"`
}

func (t *Transaction) Put(userId string) error {
	items := []dynamodb.DynamoItem{
		{
			ID:       t.ID,
			DataType: "transaction-date",
			Data:     strconv.Itoa(t.Date),
			DataKind: "transaction",
		},
		{
			ID:       t.ID,
			DataType: "transaction-food",
			Data:     t.Food.ID,
			DataKind: "transaction",
		},
		{
			ID:       t.ID,
			DataType: "transaction-user",
			Data:     userId,
		},
	}
	return dynamodb.BatchPut(items)
}
