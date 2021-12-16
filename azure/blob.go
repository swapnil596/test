package azure

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/google/uuid"
)

func GetBlobName() string {
	t := time.Now()
	uuid, _ := uuid.NewRandom()

	return fmt.Sprintf("%s-%v.txt", t.Format("20060102"), uuid)
}

func UploadBytesToBlob(b []byte) (string, error) {
	azrKey, accountName, endPoint, container := GetAccountInfo()           // This is our account info method
	u, _ := url.Parse(fmt.Sprint(endPoint, container, "/", GetBlobName())) // This uses our Blob Name Generator to create individual blob urls
	credential, err := azblob.NewSharedKeyCredential(accountName, azrKey)  // Finally we create the credentials object required by the uploader
	if err != nil {
		return "", err
	}

	// Another Azure Specific object, which combines our generated URL and credentials
	blockBlobUrl := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))

	ctx := context.Background() // We create an empty context (https://golang.org/pkg/context/#Background)

	// Provide any needed options to UploadToBlockBlobOptions (https://godoc.org/github.com/Azure/azure-storage-blob-go/azblob#UploadToBlockBlobOptions)
	o := azblob.UploadToBlockBlobOptions{
		BlobHTTPHeaders: azblob.BlobHTTPHeaders{
			ContentType: "text/plain", //  Add any needed headers here
		},
	}

	// Combine all the pieces and perform the upload using UploadBufferToBlockBlob (https://godoc.org/github.com/Azure/azure-storage-blob-go/azblob#UploadBufferToBlockBlob)
	_, err = azblob.UploadBufferToBlockBlob(ctx, b, blockBlobUrl, o)
	return blockBlobUrl.String(), err
}

func GetBlobData(blobUrl string) (string, error) {
	azrKey, accountName, _, _ := GetAccountInfo()                         // This is our account info method
	u, _ := url.Parse(blobUrl)                                            // This uses our Blob Name Generator to create individual blob urls
	credential, err := azblob.NewSharedKeyCredential(accountName, azrKey) // Finally we create the credentials object required by the uploader
	if err != nil {
		return "", err
	}

	// Another Azure Specific object, which combines our generated URL and credentials
	// blockBlobUrl := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))

	ctx := context.Background() // We create an empty context (https://golang.org/pkg/context/#Background)

	// Combine all the pieces and perform the upload using UploadBufferToBlockBlob (https://godoc.org/github.com/Azure/azure-storage-blob-go/azblob#UploadBufferToBlockBlob)
	// opts := azblob.DownloadFromBlobOptions{}
	// var b []byte

	blbUrl := azblob.NewBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))

	resp, err := blbUrl.Download(ctx, 0, 0, azblob.BlobAccessConditions{}, false, azblob.ClientProvidedKeyOptions{})
	if err != nil {
		return "", err
	}

	defer resp.Response().Body.Close()
	body, err := ioutil.ReadAll(resp.Body(azblob.RetryReaderOptions{}))

	return string(body), nil
}

func DeleteBlobData(blobUrl string) (string, error) {
	azrKey, accountName, _, _ := GetAccountInfo()                         // This is our account info method
	u, _ := url.Parse(blobUrl)                                            // This uses our Blob Name Generator to create individual blob urls
	credential, err := azblob.NewSharedKeyCredential(accountName, azrKey) // Finally we create the credentials object required by the uploader
	if err != nil {
		return "", err
	}

	// Another Azure Specific object, which combines our generated URL and credentials
	// blockBlobUrl := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))

	ctx := context.Background() // We create an empty context (https://golang.org/pkg/context/#Background)

	// Combine all the pieces and perform the upload using UploadBufferToBlockBlob (https://godoc.org/github.com/Azure/azure-storage-blob-go/azblob#UploadBufferToBlockBlob)
	// opts := azblob.DownloadFromBlobOptions{}
	// var b []byte

	blbUrl := azblob.NewBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))

	// resp, err := blbUrl.Download(ctx, 0, 0, azblob.BlobAccessConditions{}, false, azblob.ClientProvidedKeyOptions{})
	resp, err := blbUrl.Delete(ctx, azblob.DeleteSnapshotsOptionNone, azblob.BlobAccessConditions{})
	if err != nil {
		return "", err
	}

	defer resp.Response().Body.Close()
	body, err := ioutil.ReadAll(resp.Response().Body)

	return string(body), nil
}
