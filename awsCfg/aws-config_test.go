// aws-config_test.go
package awsCfg

import (
	"testing"
)

func TestAWSConfig(t *testing.T) {
	// Call the awsConfig function
	config, err := NewAwsConfig()
	if err != nil {
		t.Fatalf("Failed to create AWSConfig: %v", err)
	}

	// Verify that the AWS session is correctly created
	if config.Session == nil {
		t.Errorf("Expected Session to be non-nil")
	}

	// Verify that the DynamoDB client is correctly created
	if config.DocClient == nil {
		t.Errorf("Expected DocClient to be non-nil")
	}

	// Verify that the S3 client is correctly created
	if config.S3Client == nil {
		t.Errorf("Expected S3Client to be non-nil")
	}
}
