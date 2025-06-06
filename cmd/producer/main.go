package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	util "github.com/brainplot/sqs-example"
)

func main() {
	// Load environment variable
	queueName := os.Getenv("SQS_QUEUE_NAME")
	if queueName == "" {
		log.Fatal("Environment variable SQS_QUEUE_NAME is not set")
	}

	// Load AWS config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}

	// Create SQS client
	sqsClient := sqs.NewFromConfig(cfg)

	// Get queue URL from name
	queueURL, err := util.GetQueueURL(context.TODO(), sqsClient, queueName)
	if err != nil {
		log.Fatalf("failed to get queue URL: %v", err)
	}
	log.Printf("Sending messages to queue: %s\n", queueURL)

	// Loop indefinitely sending random numbers
	for {
		num := rand.Intn(1000) // random number between 0-999
		err := sendMessage(context.TODO(), sqsClient, queueURL, fmt.Sprintf("%d", num))
		if err != nil {
			log.Printf("error sending message: %v", err)
		} else {
			log.Printf("sent number: %d", num)
		}
		time.Sleep(2 * time.Second)
	}
}

func sendMessage(ctx context.Context, client *sqs.Client, queueURL string, message string) error {
	_, err := client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(message),
	})
	return err
}
