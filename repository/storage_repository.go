package repository

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/isd-sgcu/cutu2025-backend/utils"
)

type StorageRepository struct {
	S3Client *s3.S3
}

func NewStorageRepository(s3Client *s3.S3) *StorageRepository {
	return &StorageRepository{S3Client: s3Client}
}

func (c *StorageRepository) UploadFile(bucketName, objectKey string, buffer *bytes.Reader) (string, error) {
	// Upload the file to S3 (Google Cloud Storage in this case)
	_, err := c.S3Client.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(bucketName),
		Key:                aws.String(objectKey),
		Body:               buffer,
		ACL:                aws.String(s3.ObjectCannedACLPublicRead),
		ContentDisposition: aws.String("inline"),                   // Ensure file is displayed in the browser
		ContentType:        aws.String("application/octet-stream"), // Set MIME type for better browser handling
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}
	url := utils.GetEnv("PRODUCTION_BASE_URL", "")
	// Construct the URL of the uploaded file (using Google Cloud Storage URL format)
	fileURL := fmt.Sprintf("%s/api/users/image/%s", url, objectKey)

	// Return the URL of the uploaded file
	return fileURL, nil
}

func (c *StorageRepository) GetFileURL(bucketName, objectKey string) string {
	// Construct the URL of the uploaded file (using Google Cloud Storage URL format)
	fileURL := fmt.Sprintf("https://%s.storage.googleapis.com/%s", bucketName, objectKey)

	// Return the URL of the uploaded file
	return fileURL
}

func (c *StorageRepository) DownloadFile(bucketName, objectKey, filePath string) error {
	result, err := c.S3Client.GetObject(&s3.GetObjectInput{
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

func (c *StorageRepository) DeleteFile(bucketName, objectKey string) error {
	_, err := c.S3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return fmt.Errorf("failed to delete object, %v", err)
	}

	return nil
}
