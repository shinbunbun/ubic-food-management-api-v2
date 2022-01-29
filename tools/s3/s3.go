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

	ses := session.Must(session.NewSession())
	var s3vc *s3.S3

	if os.Getenv("AWS_SAM_LOCAL") == "true" {
		s3vc = s3.New(ses, &aws.Config{
			S3ForcePathStyle: aws.Bool(true),
			Region:           aws.String("ap-north-east-1"),
			Endpoint:         aws.String("http://localstack:4566"),
			DisableSSL:       aws.Bool(false),
		})
	} else {
		s3vc = s3.New(ses)
	}

	Uploader = s3manager.NewUploaderWithClient(s3vc)
}

func Upload(buf *bytes.Buffer, fileName string, contentType string) (string, error) {
	res, err := Uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String(contentType),
		ACL:         aws.String("public-read"),
	})
	if err != nil {
		fmt.Printf("s3 upload error: %s\n", err.Error())
		return "", err
	}
	return res.Location, nil
}
