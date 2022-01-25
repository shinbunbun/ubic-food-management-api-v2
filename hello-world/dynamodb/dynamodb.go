package dynamodb

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
)

var table dynamo.Table

func CreateTable() {
	disableSsl := false

	dynamoDBEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	if len(dynamoDBEndpoint) != 0 {
		disableSsl = true
	}

	dynamoDBRegion := "ap-north-east-1"

	ses := session.Must(session.NewSession())

	db := dynamo.New(ses, &aws.Config{
		Region:     aws.String(dynamoDBRegion),
		Endpoint:   aws.String(dynamoDBEndpoint),
		DisableSSL: aws.Bool(disableSsl),
	})

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

func GetByDataKind(dataKind string) (DynamoItem, error) {
	var readResult DynamoItem
	err := table.Get("DataKind", dataKind).Index("DataKind-index").One(&readResult)
	if err != nil {
		fmt.Printf("Failed to get item[%v]\n", err)
		return DynamoItem{}, err
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
	wrote, err := table.Batch().Write().Put(items).Run()
	if wrote != len(items) {
		return fmt.Errorf("Failed to put %d items", len(items))
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
