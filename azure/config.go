package azure

import "fmt"

func GetAccountInfo() (string, string, string, string) {
	azrKey := "uGFxW6GfRg5Zg6hlHxDPARTpcy4k3AJv0ibn+qmGh6ZUICQYYnRo2et+kiUqAmo48/cYlYSIOjwjC7Mxk5VwEw=="
	azrBlobAccountName := "csg10032001a70f2ae0"
	azrPrimaryBlobServiceEndpoint := fmt.Sprintf("https://%s.blob.core.windows.net/", azrBlobAccountName)
	azrBlobContainer := "flowxpert-blobs"

	return azrKey, azrBlobAccountName, azrPrimaryBlobServiceEndpoint, azrBlobContainer
}
