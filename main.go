package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	awsCfg, err := awsConfig() //awsConfig func from aws-config.go
	if err != nil {
		log.Fatalf("Failed to set up AWS session: %v", err)
	}

	fmt.Println("DynamoDB and S3 clients created successfully")

	awsCfg.ListBuckets()
}

// ListBuckets lists the S3 buckets
func (cfg *AWSConfig) ListBuckets() {
	result, err := cfg.S3Client.ListBuckets(nil)
	if err != nil {
		log.Fatalf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")
	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
