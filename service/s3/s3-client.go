package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

type Client struct {
	Service *s3.S3
}

func NewS3Client() *Client {
	region := os.Getenv("AWS_REGION")
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("AWS_SECRET_ACCESS_KEY")

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, accessKeySecret, ""),
	}))

	return &Client{
		Service: s3.New(sess),
	}
}
