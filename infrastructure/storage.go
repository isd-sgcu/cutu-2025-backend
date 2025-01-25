package infrastructure

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/isd-sgcu/cutu2025-backend/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	client *s3.S3
}

// NewS3Client initializes a new S3 client using AWS SDK v1
func NewS3Client(cfg *config.Config) *S3Client {
	// Get credentials from environment variables
	accessKey := cfg.AWSAccessKeyID
	secretKey := cfg.AWSSecretAccessKey
	region := cfg.AWSRegion

	if accessKey == "" || secretKey == "" || region == "" {
		panic("AWS credentials or region not found in .env file")
	}

	// Create the AWS credentials object
	creds := credentials.NewStaticCredentials(accessKey, secretKey, "")

	// Create a new session with AWS SDK v1
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(region),
		Credentials:      creds,
		Endpoint:         aws.String("https://storage.googleapis.com"),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to create session: %v", err))
	}

	// Create the S3 client
	client := s3.New(sess)

	// Check if the client is successfully connected by listing the buckets
	_, err = client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to S3 service: %v", err))
	}

	fmt.Println("Successfully connected to the S3 service")

	return &S3Client{
		client: client,
	}
}

func (c *S3Client) UploadFile(bucketName, objectKey string, buffer *bytes.Reader) (string, error) {
	// Upload the file to S3 (Google Cloud Storage in this case)
	_, err := c.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   buffer,
		ACL:    aws.String(s3.ObjectCannedACLPublicRead),
		ContentDisposition: aws.String("inline"), // Ensure file is displayed in the browser
		ContentType:        aws.String("application/octet-stream"), // Set MIME type for better browser handling
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	// Construct the URL of the uploaded file (using Google Cloud Storage URL format)
	fileURL := fmt.Sprintf("https://%s.storage.googleapis.com/%s", bucketName, objectKey)

	// Return the URL of the uploaded file
	return fileURL, nil
}


func (c *S3Client) DownloadFile(bucketName, objectKey, filePath string) error {
	result, err := c.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return fmt.Errorf("failed to get object, %v", err)
	}
	defer result.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s, %v", filePath, err)
	}
	defer file.Close()

	_, err = io.Copy(file, result.Body)
	if err != nil {
		return fmt.Errorf("failed to write file, %v", err)
	}

	return nil
}

func (c *S3Client) DeleteFile(bucketName, objectKey string) error {
	_, err := c.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return fmt.Errorf("failed to delete object, %v", err)
	}

	return nil
}
