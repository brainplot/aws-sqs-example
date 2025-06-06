package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	util "github.com/brainplot/sqs-example"
)

func main() {
	// Load queue name from environment
	queueName := os.Getenv("SQS_QUEUE_NAME")
	if queueName == "" {
		log.Fatal("Environment variable SQS_QUEUE_NAME is not set")
	}

	// Load AWS config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	sqsClient := sqs.NewFromConfig(cfg)

	// Resolve queue URL
	queueURL, err := util.GetQueueURL(context.TODO(), sqsClient, queueName)
	if err != nil {
		log.Fatalf("failed to get queue URL: %v", err)
	}
	log.Printf("Listening to queue: %s", queueURL)

	// Poll every 5 seconds
	for {
		err := receiveMessages(context.TODO(), sqsClient, queueURL)
		if err != nil {
			log.Printf("error receiving messages: %v", err)
		}
		time.Sleep(5 * time.Second)
	}
}

func receiveMessages(ctx context.Context, client *sqs.Client, queueURL string) error {
	out, err := client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: 5,  // Adjust based on expected volume
		WaitTimeSeconds:     10, // Long polling
		VisibilityTimeout:   30, // Optional: time before message becomes visible again
	})
	if err != nil {
		return err
	}

	if len(out.Messages) == 0 {
		log.Println("No messages received")
		return nil
	}

	for _, msg := range out.Messages {
		fmt.Printf("Received: %s\n", *msg.Body)

		// Delete the message after processing
		_, err := client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(queueURL),
			ReceiptHandle: msg.ReceiptHandle,
		})
		if err != nil {
			log.Printf("Failed to delete message: %v", err)
		} else {
			log.Printf("Deleted message ID: %s", *msg.MessageId)
		}
	}

	return nil
}
