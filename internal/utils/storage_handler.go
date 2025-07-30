package utils

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"time"

	"firebase.google.com/go/v4/storage"
)

// FIREBASE STORAGE
func UploadFileToFirebase(sc *storage.Client, file *multipart.FileHeader, folderName string) (string, error) {
	ctx := context.Background()

	// Open the file from multipart.FileHeader
	fileContent, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer fileContent.Close()

	// Get the default bucket
	bucket, err := sc.DefaultBucket()
	if err != nil {
		return "", fmt.Errorf("failed to get default bucket: %v", err)
	}

	// Generate a unique object name
	objectName := fmt.Sprintf("%s/%d%s", folderName, time.Now().UnixNano(), filepath.Ext(file.Filename))
	contentType := file.Header.Get("Content-Type")

	// Create a writer for the object
	writer := bucket.Object(objectName).NewWriter(ctx)
	writer.ContentType = contentType

	// Copy the file content to the writer
	if _, err := io.Copy(writer, fileContent); err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	// Close the writer to finalize the upload
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	// attrs, err := bucket.Object(objectName).Attrs(context.Background())
	// if err != nil {
	// 	return "", fmt.Errorf("failed to get file attributes: %v", err)
	// }

	// Return the MediaLink (public URL if the file is publicly accessible)
	return objectName, nil
}

func GetFirebasePublicUrl(objectName string) string {
	encodedObjectName := url.QueryEscape(objectName)
	basePath := "https://firebasestorage.googleapis.com/v0/b/test-sk-web-dev.firebasestorage.app/o/"

	return basePath + encodedObjectName + "?alt=media"
}
