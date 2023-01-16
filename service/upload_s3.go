package service

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/faridlan/go-s3/helper"
	"github.com/faridlan/go-s3/model"
)

func UploadS3(body io.Reader) model.Upload {
	helper.LoadEnv()
	awsAccessKeyID := helper.GetEnvWithKey("AWS_ACCESS_KEY_ID")
	fmt.Println("My access key ID is ", awsAccessKeyID)

	sess := session.Must(session.NewSession())

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	s := helper.RandStringRunes(10)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(helper.GetEnvWithKey("BUCKET_NAME")),
		Key:         aws.String("test/" + s + ".png"),
		Body:        body,
		ContentType: aws.String("image/png"),
		ACL:         aws.String("public-read"),
	})

	if err != nil {
		panic(err)
	}

	return model.Upload{
		ImageUrl: aws.StringValue(&result.Location),
	}
	// fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
}
