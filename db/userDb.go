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

func (dm UserDbManger) AddUser(u User) error {
	dc := dm.Session()
	item, error := dynamodbattribute.MarshalMap(u)
	if error != nil {
		fmt.Println("error in marshal")
		return error
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("my_table"),
	}
	_, err := dc.PutItem(input)
	return err
}

func (dm UserDbManger) ListUser() ([]User, error) {
	dc := dm.Session()
	input := &dynamodb.ScanInput{
		TableName: aws.String("my_table"),
	}
	result, err := dc.Scan(input)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	u := []User{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &u)
	return u, err
}

func (dm UserDbManger) UpdateUser(u User) error {
	dc := dm.Session()
	// item, error := dynamodbattribute.MarshalMap(u)

	type UpdateInfo struct {
		Age    int    `json:":age"`
		Gender string `json:":gender"`
	}
	up, err := dynamodbattribute.MarshalMap(UpdateInfo{
		Age:    u.Age,
		Gender: u.Gender,
	})

	if err != nil {
		fmt.Println("Something went wrong")
	}

	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"name": {S: aws.String(u.Name)},
		},
		TableName:                 aws.String("my_table"),
		UpdateExpression:          aws.String("set age = :age, gender = :gender"),
		ExpressionAttributeValues: up,
		ReturnValues:              aws.String("NONE"),
	}

	_, err1 := dc.UpdateItem(input)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	return err1
}
