package awscourier

// Credentials store AWS credentials
type Credentials struct {
	AwsAccessKeyId     *string `json:"aws_access_key_id,omitempty"`
	AwsSecretAccessKey *string `json:"aws_secret_access_key,omitempty"`
	AwsSessionToken    *string `json:"aws_session_token,omitempty"`
}

// CredentialsService defines a service to return AWS credentials
type CredentialsService interface {
	Credentials(profile *string) (*Credentials, error)
}
