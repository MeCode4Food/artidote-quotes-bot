package util

import (
	"artidote-quote/constants"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// GetRoomIDsFromDB gets list of room IDs from dynamoDB
func GetRoomIDsFromDB(session *session.Session) []string {

	svc := dynamodb.New(session)
	println("getting item 2")
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
