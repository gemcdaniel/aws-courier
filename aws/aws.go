package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/gemcdaniel/aws-courier"
)

// CredentialsService is an AWS implementation fo the CredentialsService
type CredentialsService struct {
	AwsSharedCredentialFile string
}

// NewCredentialsService create a new CredentialsService struct
func NewCredentialsService(filename string) *CredentialsService {
	return &CredentialsService{
		AwsSharedCredentialFile: filename,
	}
}

// Credentials is an implementation fo the CredentialsService interface. It returns the credentials provided by AWS from the local filesystem
func (cs *CredentialsService) Credentials(profile *string) (*awscourier.Credentials, error) {
	fileCredentials := credentials.NewSharedCredentials(cs.AwsSharedCredentialFile, aws.StringValue(profile))
	creds, err := fileCredentials.Get()
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve credentials for '%s' profile: %v", aws.StringValue(profile), err)
	}

	awsCredentials := &awscourier.Credentials{}

	if creds.AccessKeyID != "" {
		awsCredentials.AwsAccessKeyId = aws.String(creds.AccessKeyID)
	}

	if creds.SecretAccessKey != "" {
		awsCredentials.AwsSecretAccessKey = aws.String(creds.SecretAccessKey)
	}

	if creds.SessionToken != "" {
		awsCredentials.AwsSessionToken = aws.String(creds.SessionToken)
	}

	return awsCredentials, nil
}
