package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"os"
)

type Client struct {
	Service *s3.S3
}

func NewS3Client() *Client {
	keyId := os.Getenv("S3_KEY_ID")
	secretKey := os.Getenv("S3_SECRET_KEY")
	region := os.Getenv("S3_REGION")
	endPoint := os.Getenv("S3_ENDPOINT")

	sess, err := session.NewSession(&aws.Config{
		Endpoint:    aws.String(endPoint),
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(keyId, secretKey, ""),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %s", err)
	}

	return &Client{
		Service: s3.New(sess),
	}
}
