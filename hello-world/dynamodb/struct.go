package dynamodb

type DynamoItem struct {
	ID       string
	DataType string
	Data     string `dynamo:",omitempty"`
	DataKind string
	IntData  int `dynamo:",omitempty"`
}
