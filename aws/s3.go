package aws

// import (
// 	"api-registration-backend/config"
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"reflect"
// 	"strings"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/awserr"
// 	"github.com/aws/aws-sdk-go/service/s3"
// )

// func GetS3Object(key string) (interface{}, error) {

// 	// The session the S3 will use
// 	svc := GetS3Client()

// 	// load configurations
// 	conf := config.GetConfigurations()

// 	params := &s3.GetObjectInput{
// 		Bucket: aws.String(conf.GetString("aws.s3.bucket")),
// 		Key:    aws.String(key),
// 	}

// 	resp, err := svc.GetObject(params)
// 	if err != nil {
// 		if aerr, ok := err.(awserr.Error); ok {
// 			switch aerr.Code() {
// 			case s3.ErrCodeNoSuchKey:
// 				log.Println(s3.ErrCodeNoSuchKey, aerr.Error())
// 			default:
// 				log.Printf("S3 Object access error, %s\n", aerr.Error())
// 			}
// 		} else {
// 			// Print the error, cast err to awserr.Error to get the Code and
// 			// Message from an error.
// 			log.Println(err.Error())
// 		}
// 		return nil, err
// 	}

// 	// to avoid resource leak, closing the response body once finished
// 	defer func(Body io.ReadCloser) {
// 		err := Body.Close()
// 		if err != nil {

// 		}
// 	}(resp.Body)

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Printf("Error reading s3 file, %s\n", err.Error())
// 	}

// 	bodyString := fmt.Sprintf("%s", body)

// 	var s3data map[string]interface{}
// 	decoder := json.NewDecoder(strings.NewReader(bodyString))
// 	err = decoder.Decode(&s3data)
// 	if err != nil {
// 		log.Println("Error decoding file")
// 	}

// 	fmt.Println(reflect.TypeOf(s3data))

// 	return s3data, nil
// }

// func UploadFileToS3(key string, reader io.Reader, size int64) (string, error) {

// 	// The session the S3 will use
// 	svc := GetS3Client()

// 	// load configurations
// 	conf := config.GetConfigurations()

// 	// read the file content into a buffer
// 	buffer := make([]byte, size)

// 	_, err := reader.Read(buffer)
// 	if err != nil {
// 		return "", err
// 	}

// 	resp, err := svc.PutObject(&s3.PutObjectInput{
// 		Bucket:               aws.String(conf.GetString("aws.s3.bucket")),
// 		Key:                  aws.String(key),
// 		ACL:                  aws.String("private"),
// 		Body:                 bytes.NewReader(buffer),
// 		ContentLength:        aws.Int64(size),
// 		ContentType:          aws.String(http.DetectContentType(buffer)),
// 		ContentDisposition:   aws.String("attachment"),
// 		ServerSideEncryption: aws.String("AES256"),
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	log.Println(resp)
// 	return key, err
// }
