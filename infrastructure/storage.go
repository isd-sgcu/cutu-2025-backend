package infrastructure

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/isd-sgcu/cutu2025-backend/config"
)

// ConnectToS3 initializes a new S3 client using AWS SDK v1
func ConnectToS3(cfg *config.Config) *s3.S3 {
	// ตรวจสอบว่าข้อมูล AWS credentials และ region ถูกตั้งค่าใน config หรือไม่
	accessKey := cfg.AWSAccessKeyID
	secretKey := cfg.AWSSecretAccessKey
	region := cfg.AWSRegion

	if accessKey == "" || secretKey == "" || region == "" {
		log.Fatalf("AWS credentials or region not found in configuration: AccessKey=%s, Region=%s", accessKey, region)
	}

	// สร้าง AWS credentials object
	creds := credentials.NewStaticCredentials(accessKey, secretKey, "")

	// สร้าง session ใหม่
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(region),
		Credentials:      creds,
		Endpoint:         aws.String("https://storage.googleapis.com"),
		S3ForcePathStyle: aws.Bool(true),                              
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	// สร้าง S3 client
	client := s3.New(sess)

	// ตรวจสอบการเชื่อมต่อกับ S3 โดยการเรียก ListBuckets
	_, err = client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("Failed to connect to S3 service: %v", err)
	}

	log.Println("Successfully connected to the S3 service")
	return client
}
