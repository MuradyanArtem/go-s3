package client

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3ClientConfig configuration struct for s3 instance
type S3ClientConfig struct {
	Address string
	Access  string
	Secret  string
	Token   string
	Region  string
}

// OpenS3 return instance of s3 client
func OpenS3(config S3ClientConfig) (*s3.S3, error) {
	session, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			config.Access,
			config.Secret,
			config.Token,
		),
		Endpoint:         aws.String(config.Address),
		Region:           aws.String(config.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	})

	if err != nil {
		return nil, fmt.Errorf("open s3 failed: %w", err)
	}

	return s3.New(session), nil
}
