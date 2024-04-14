package src

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
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

func (config AWSConfig) ListForQueue() []string {

	var queueUrls []string
	pages := sqs.NewListQueuesPaginator(config.q, &sqs.ListQueuesInput{})

	for pages.HasMorePages() {
		output, err := pages.NextPage(context.TODO())

		if err != nil {
			log.Printf("Could't get queue, Error : %v\n", err)
			break
		} else {
			queueUrls = append(queueUrls, output.QueueUrls...)
		}
	}

	return queueUrls
}

func (config AWSConfig) RetrieveQueue(queueName string) (string, error) {

	urls := config.ListForQueue()
	for _, url := range urls {
		if strings.Contains(url, queueName) {
			return url, nil
		}
	}

	return "", fmt.Errorf("not Exists %s", queueName)
}

func (config AWSConfig) RetrieveQueueProperty(queueName string) (string, error) {

	url, err := config.RetrieveQueue(queueName)
	if err != nil {
		return "", err
	}

	attrName := types.QueueAttributeNameQueueArn
	attr, err := config.q.GetQueueAttributes(context.TODO(), &sqs.GetQueueAttributesInput{
		QueueUrl:       aws.String(url),
		AttributeNames: []types.QueueAttributeName{attrName},
	})

	if err != nil {
		return "", err
	}

	return attr.Attributes[string(attrName)], nil
}

// types.MessageAttributeValue
func (config AWSConfig) SendMessage(messageBody, queueURL string, attr map[string]types.MessageAttributeValue) error {

	_, err := config.q.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:          &queueURL,
		MessageBody:       &messageBody,
		MessageAttributes: attr,
	})

	return err
}

func (config AWSConfig) SendBatchMessage(queueUrl string, attr []types.SendMessageBatchRequestEntry) error {

	req := &sqs.SendMessageBatchInput{
		QueueUrl: &queueUrl,
		Entries:  attr,
	}

	_, err := config.q.SendMessageBatch(context.TODO(), req)

	return err
}

func (config AWSConfig) pollMessage() {

}

func (config AWSConfig) CreateQueue() {
	// Terraform 으로 대체
	// Queue, FIFO_Queue, DeadLetter Queue
}
