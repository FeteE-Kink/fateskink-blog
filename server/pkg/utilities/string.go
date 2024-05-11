package utilities

import (
	"fmt"
	"os"
)

func GetAvatarUrl(key string) *string {
	var url string

	if key == "" || &key == nil {
		return nil
	}

	if os.Getenv("UPLOAD_LOCALLY_PATH") != "" {
		url = os.Getenv("UPLOAD_LOCALLY_PATH") + key
	} else {
		bucketName := os.Getenv("GCS_BUCKET_NAME")
		url = fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, key)
	}

	return &url
}
