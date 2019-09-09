package db

import (
	"fmt"
	. "lyang/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/graniticio/granitic/logging"
)

type UserDynamoDBClient struct {
	Svc dynamodbiface.DynamoDBAPI
}

type UserDbManger struct {
	Log logging.Logger
}

func (dm UserDbManger) Session() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	dbSvc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://127.0.0.1:8000")})
	return dbSvc
}

func (dm UserDbManger) AddUser(w User) error {
	dc := dm.Session()
	item, error := dynamodbattribute.MarshalMap(w)
	if error != nil {
		fmt.Println("error in marshal")
		return error
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("my_table"),
	}
	fmt.Println(input)

	_, error1 := dc.PutItem(input)
	return error1
}
