package s3

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var Uploader *s3manager.Uploader

func init() {
	var endpoint string
	disableSsl := false

	if os.Getenv("AWS_SAM_LOCAL") == "true" {
		endpoint = "http://localstack:4566"
		disableSsl = true
	}
	/* ses, err := session.NewSession(&aws.Config{
		Region:     aws.String(endpoints.ApNortheast1RegionID),
		Endpoint:   aws.String(endpoint),
		DisableSSL: aws.Bool(disableSsl),
	}) */
	ses := session.Must(session.NewSession())
	s3vc := s3.New(ses, &aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String("ap-north-east-1"),
		Endpoint:         aws.String(endpoint),
		DisableSSL:       aws.Bool(disableSsl),
	})

	Uploader = s3manager.NewUploaderWithClient(s3vc)
}

func Upload(buf *bytes.Buffer, fileName string, contentType string) (string, error) {
	res, err := Uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		fmt.Printf("s3 upload error: %s\n", err.Error())
		return "", err
	}
	return res.Location, nil
}
