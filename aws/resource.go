package aws

// import (
// 	"api-registration-backend/config"
// 	"log"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/credentials"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/dynamodb"
// 	"github.com/aws/aws-sdk-go/service/s3"
// )

// var dynamo *dynamodb.DynamoDB
// var svc *s3.S3

// func Init() {
// 	// load configurations
// 	c := config.GetConfigurations()

// 	// Initializing a session that the SDK will use to load
// 	// credentials from the shared file ~/.aws/credentials

// 	// For some weird reason golang cannot access ~/.aws/credentials
// 	creds := credentials.NewEnvCredentials()

// 	// Retrieve the credentials value
// 	_, err := creds.Get()
// 	if err != nil {
// 		// application should not start if credentials are not accessible
// 		// log.Fatalf("aws credentials error >> %v\n", err)
// 	}

// 	sess, err := session.NewSession(&aws.Config{
// 		Region:      aws.String(c.GetString("aws.region")),
// 		Credentials: credentials.NewStaticCredentials(c.GetString("aws.key"), c.GetString("aws.secret"), ""),
// 	})

// 	if err != nil {
// 		log.Fatalf("AWS session error, %v", err)
// 	}

// 	// creating a dynamo service client
// 	dynamo = dynamodb.New(sess)

// 	// creating a s3 service client
// 	svc = s3.New(sess)
// }

// func GetDynamoClient(tableName string, primaryKey string, sortKey string) DBConfig {

// 	return DBConfig{
// 		dbService:  dynamo,
// 		primaryKey: primaryKey,
// 		sortKey:    sortKey,
// 		tableName:  tableName,
// 	}
// }

// func GetS3Client() *s3.S3 {
// 	return svc
// }
