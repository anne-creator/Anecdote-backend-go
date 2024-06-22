// Description: This file contains the AWSConfig struct and the awsConfig function.
// @returns: AWSConfig struct{} a session and service clients for DynamoDB and S3}
package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	TableName    = "AnecdoteTask"
	S3BucketName = "anecdote-web"
)

// AWSConfig TYPE holds the AWS session and service clients
type AWSConfig struct {
	Session   *session.Session
	DocClient *dynamodb.DynamoDB
	S3Client  *s3.S3
}

func awsConfig() (*AWSConfig, error) {
	// // Load environment variables
	// accessKeyID := os.Getenv("aws_access_key_id")
	// secretAccessKey := os.Getenv("aws_secret_access_key")

	// if accessKeyID == "" || secretAccessKey == "" {
	// 	log.Fatalf("Environment variables MY_AWS_ACCESS_KEY and MY_AWS_SECRET_KEY must be set")
	// }

	// fmt.Println(accessKeyID)

	// Create a new session using the specified credentials and region
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	fmt.Println("docClient and S3Client work here")

	return &AWSConfig{
		Session:   sess,
		DocClient: dynamodb.New(sess),
		S3Client:  s3.New(sess),
	}, nil
}
