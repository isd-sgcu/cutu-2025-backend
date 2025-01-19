package infrastructure

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

func UploadFileToS3(filePath, key string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	// Example of uploading a file; you need to customize it for your use case
	// Here is a simple example of uploading a file to S3
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	bucketName := os.Getenv("S3_BUCKET_NAME")
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &key,
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", os.Getenv("S3_BUCKET_NAME"), key), nil
}
