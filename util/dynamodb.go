package util

import (
	"artidote-quote/constants"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// GetRoomIDsFromDB gets list of room IDs from dynamoDB
func GetRoomIDsFromDB(session *session.Session) []string {

	svc := dynamodb.New(session)

	result, err := svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(constants.DBName),
		KeyConditions: map[string]*dynamodb.Condition{
			"DocID": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{S: aws.String(constants.DocID.ChatIDList)},
				},
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	if len(result.Items) == 0 {
		log.Fatalln("No results returned")
	}

	type DynamoDBItem struct {
		DocID       string
		CreatedDate string
		DocName     string
		IDList      []string
	}

	item := DynamoDBItem{}

	err = dynamodbattribute.UnmarshalMap(result.Items[0], &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return item.IDList
}
