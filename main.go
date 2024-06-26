package main

import (
	"fmt"
	"log"

	"github.com/anne-creator/Anecdote-backend-go/awsCfg"
)

func main() {
	AWSConfig, err := awsCfg.NewAwsConfig() //from calling func in awsCfg to create awsCOnfig variable
	if err != nil {
		log.Fatalf("Failed to set up AWS session: %v", err)
	}

	fmt.Println("DynamoDB and S3 clients created successfully")

	AWSConfig.ListBuckets() //
}
