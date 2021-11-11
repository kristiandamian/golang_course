package dynamoinstance

//https://github.com/aws/aws-sdk-go-v2
import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

func Save() (string, error) {
	svc, err := getDynamoClient()
	if err != nil {
		return "", err
	}
	id := generateUuid()

	input := &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: id},
		},
		TableName: aws.String("MyTable"),
	}

	_, err = svc.PutItem(context.TODO(), input)
	if err != nil {
		return "", err
	}
	fmt.Println("Successfully added '" + id + "  to MyTable ")
	return id, nil
}

func Read(id string) error {
	svc, err := getDynamoClient()
	if err != nil {
		return err
	}

	result, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("MyTable"),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return err
	}
	if result.Item == nil {
		return errors.New("item not found")
	}
	fmt.Println("Successfully readed the item with the Id: '" + id + "  in MyTable ")
	return err
}

func GetTablesNames() error {
	svc, err := getDynamoClient()
	if err != nil {
		return err
	}
	resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})

	if err != nil {
		return err
	}

	fmt.Println("Tables:")
	for _, tableName := range resp.TableNames {
		fmt.Println(tableName)
	}

	return nil
}

func getDynamoClient() (*dynamodb.Client, error) {
	cfg, err := getCredentials()
	if err != nil {
		return nil, err
	}
	svc := dynamodb.NewFromConfig(*cfg)
	return svc, err
}

func getCredentials() (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func generateUuid() string {
	return uuid.New().String()
}
