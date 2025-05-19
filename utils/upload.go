package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// UploadToS3 uploads a file to the specified S3 bucket
func UploadToS3(file []byte, filename string) error {
	awsRegion := "us-east-1"
	bucket := "your-aidvault-bucket-name"

	// Load credentials from environment variables or config
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})
	if err != nil {
		log.Println("Failed to create session:", err)
		return err
	}

	uploader := s3.New(sess)

	// Upload input params
	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(filepath.Base(filename)),
		Body:        bytes.NewReader(file),
		ContentType: aws.String("application/octet-stream"),
	}

	_, err = uploader.PutObject(input)
	if err != nil {
		log.Println("Failed to upload file:", err)
		return err
	}

	fmt.Println("File uploaded to S3:", filename)
	return nil
}
