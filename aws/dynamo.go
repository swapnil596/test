package aws

import (
	"api-registration-backend/models"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DBConfig struct {
	dbService  *dynamodb.DynamoDB
	primaryKey string
	sortKey    string
	tableName  string
}

func (database DBConfig) PutItem(props interface{}) (interface{}, error) {
	av, err := dynamodbattribute.MarshalMap(props)
	if err != nil {
		log.Printf("error while marshalling new property item, %v\n", err)
	}

	input := &dynamodb.PutItemInput{
		Item:                av,
		TableName:           aws.String(database.tableName),
		ConditionExpression: aws.String("attribute_not_exist(id)"),
	}

	_, err = database.dbService.PutItem(input)
	if err != nil {
		log.Printf("error while storing item, %v", err)
	}

	return props, err
}

func (database DBConfig) UpdateItem(api models.API, updateExpression string, expr map[string]*dynamodb.AttributeValue) error {

	key, err := dynamodbattribute.MarshalMap(api)
	if err != nil {
		log.Printf("error while marshalling old property item, %v\n", err)
		return err
	}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: expr,
		TableName:                 aws.String(database.tableName),
		Key:                       key,
		UpdateExpression:          aws.String(updateExpression),
	}

	_, err = database.dbService.UpdateItem(input)
	if err != nil {
		log.Printf("error while updating item, %v", err)
	}

	return err
}

// GetItem Returns http status code & error message if any
func (database DBConfig) GetItem(primaryKey string, sortKey string, data interface{}) (int, error) {
	av := map[string]*dynamodb.AttributeValue{
		database.primaryKey: {
			S: aws.String(primaryKey),
		},
	}

	if sortKey != "" {
		av[database.sortKey] = &dynamodb.AttributeValue{
			S: aws.String(sortKey),
		}
	}

	result, err := database.dbService.GetItem(&dynamodb.GetItemInput{
		Key:       av,
		TableName: aws.String(database.tableName),
	})

	if err != nil {
		log.Printf("data not found in dynamo, %v\n", err)
		return http.StatusNotFound, err
	}

	// test
	log.Println(result)

	err = dynamodbattribute.UnmarshalMap(result.Item, data)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, err
}
