// Description: This file contains the AWSConfig struct and the awsConfig function.
// @returns:
// AWSConfig struct{
// a session
// DynamoDB client
// S3}
package awsCfg

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

func NewAwsConfig() (*AWSConfig, error) {
	//the default credentials and region was in .aws/credentials

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

// ListBuckets lists the S3 buckets
func (cfg *AWSConfig) ListBuckets() { //
	result, err := cfg.S3Client.ListBuckets(nil)
	if err != nil {
		log.Fatalf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")
	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
