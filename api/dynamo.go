package api

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type dynamoClient interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

type dynStore struct {
	client dynamoClient
	v      string
}

type Key = map[string]types.AttributeValue

var (
	// ensure `dynStore` implements `IStore` interface
	_ Storer = &dynStore{}
)

func NewStore(ctx context.Context, version string) (*dynStore, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	return &dynStore{
		client: dynamodb.NewFromConfig(cfg),
		v:      version,
	}, nil
}

func (s *dynStore) version() string {
	if s.v == "" {
		return "dev"
	}
	return s.v
}
