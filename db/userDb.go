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
	fmt.Println("ho")
	config := &aws.Config{
		Region:   aws.String("us-west-1"),
		Endpoint: aws.String("http://localhost:8000"),
	}

	sess := session.Must(session.NewSession(config))
	svc := dynamodb.New(sess)

	return svc
}

func (dm UserDbManger) AddUser(w User) error {
	tbName := aws.String("LyangTable")
	dc := dm.Session()
	fmt.Println(*dc)
	fmt.Println("h000o")
	item, err := dynamodbattribute.MarshalMap(w)
	fmt.Println(item)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: tbName,
	}
	fmt.Println(input)

	_, err = dc.PutItem(input)
	if err != nil {
		fmt.Println("errorrrrrrrrrrr")
	}
	return err
}

func (dm UserDbManger) Test() {
	fmt.Println("inside db")
}
