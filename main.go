package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Function to generate a Presigned URL
func generatePresignedURL(minioClient *minio.Client, bucketName string, objectName string) (*url.URL, error) {
	// Properly encode the object name to match the MinIO server's expectation
	encodedObjectName := url.PathEscape(objectName)

	reqParams := make(url.Values)
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, encodedObjectName, 3600*time.Second, reqParams)
	if err != nil {
		return nil, err
	}

	return presignedURL, nil
}

func main() {
	// Initialize MinIO client
	minioClient, err := minio.New("your-minio-server:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("YOUR_ACCESS_KEY", "YOUR_SECRET_KEY", ""),
		Secure: false,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	bucketName := "your-bucket"
	objectName := "photos/my  holiday  photo.jpg"

	presignedURL, err := generatePresignedURL(minioClient, bucketName, objectName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Presigned URL:", presignedURL)
}