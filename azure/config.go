package azure

import "fmt"

func GetAccountInfo() (string, string, string, string) {
	azrKey := "ViTt18iPUaTSwdXr2NFifhxzxiDuYh/3xHyyL/O5GcZy/FwXfuGT6IxVhyy/5m6FPB5pNS5bqRehKzp7kW4qvA=="
	azrBlobAccountName := "algo360test"
	azrPrimaryBlobServiceEndpoint := fmt.Sprintf("https://%s.blob.core.windows.net/", azrBlobAccountName)
	azrBlobContainer := "flowxpert-blobs"

	return azrKey, azrBlobAccountName, azrPrimaryBlobServiceEndpoint, azrBlobContainer
}
