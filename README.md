## Instructions

### Start DynamoDB

```sh
docker run -p 8000:8000 amazon/dynamodb-local -jar DynamoDBLocal.jar -inMemory -sharedDb
```

### Create table

```sh
 aws dynamodb create-table \
    --table-name my_table \
    --attribute-definitions AttributeName=name,AttributeType=S \
    --key-schema AttributeName=name,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --endpoint-url http://localhost:8000
```

### Build and start server

```
grnc-yaml-bind && go run main.go
```