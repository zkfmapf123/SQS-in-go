package src

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type AWSConfig struct {
	q *sqs.Client
}

func New(region string) AWSConfig {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		panic(err)
	}

	return AWSConfig{
		q: sqs.NewFromConfig(cfg),
	}
}

func (config AWSConfig) CreateQueue() {
	// Terraform 으로 대체
}
