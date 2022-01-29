package dynamodb

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
)

var table dynamo.Table

func init() {

	var db *dynamo.DB
	ses := session.Must(session.NewSession())

	if os.Getenv("AWS_SAM_LOCAL") == "true" {
		db = dynamo.New(ses, &aws.Config{
			Region:      aws.String("ap-north-east-1"),
			Endpoint:    aws.String("http://dynamodb-local:8000"),
			DisableSSL:  aws.Bool(true),
			Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
		})
	} else {
		db = dynamo.New(ses)
	}

	table = db.Table("UBIC-FOOD")
}

func Put(item DynamoItem) error {
	return table.Put(item).Run()
}

func GetByIDDataType(id string, dataType string) (DynamoItem, error) {
	var readResult DynamoItem
	err := table.Get("ID", id).Range("DataType", dynamo.Equal, dataType).One(&readResult)
	if err != nil {
		fmt.Printf("Failed to get item[%v]\n", err)
		return DynamoItem{}, err
	}
	return readResult, nil
}

func GetByDataDataType(Data string, dataType string) ([]DynamoItem, error) {
	var readResult []DynamoItem
	fmt.Printf("Table: %+v\n", table)
	err := table.Get("Data", Data).Range("DataType", dynamo.Equal, dataType).Index("Data-DataType-index").All(&readResult)
	fmt.Printf("readResult: %+v\n", readResult)
	if err != nil {
		fmt.Printf("Failed to get item[%v]\n", err)
		return []DynamoItem{}, err
	}
	return readResult, nil
}

func GetByDataKind(dataKind string) ([]DynamoItem, error) {
	var readResult []DynamoItem
	err := table.Get("DataKind", dataKind).Index("DataKind-index").All(&readResult)
	if err != nil {
		fmt.Printf("Failed to get item[%v]\n", err)
		return []DynamoItem{}, err
	}
	return readResult, nil
}

func DeleteByID(id string, dataType string) error {
	return table.Delete("ID", id).Range("DataType", dataType).Run()
}

func BatchDelete(keys []dynamo.Keyed) error {
	wrote, err := table.Batch("ID", "DataType").Write().Delete(keys...).Run()
	if wrote != len(keys) {
		return fmt.Errorf("Failed to delete %d items", len(keys))
	}
	return err
}

func BatchPut(items []DynamoItem) error {

	items2 := make([]interface{}, len(items))
	for i, v := range items {
		items2[i] = v
	}

	wrote, err := table.Batch().Write().Put(items2...).Run()
	if wrote != len(items) {
		return fmt.Errorf("unexpected wrote: %d ≠ %d", wrote, len(items))
	}
	return err
}

func GetByID(id string) ([]DynamoItem, error) {
	var readResult []DynamoItem
	err := table.Get("ID", id).All(&readResult)
	if err != nil {
		fmt.Printf("Failed to get item[%v]\n", err)
		return []DynamoItem{}, err
	}
	return readResult, nil
}

func AddIntData(count int, id string, dataType string) error {
	return table.Update("ID", id).Range("DataType", dataType).Add("IntData", count).Run()
}

func GenerateID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return u.String(), nil
}
