package awsgo

import "github.com/aws/aws-sdk-go/aws"

// ConfigOptions contains the options that deliver configuration details.
type ConfigOptions interface {
	GetDefaults(ServiceType) *aws.Config
}

// GetDefault returns the default value for the given key (TODO).
func GetDefault(dir, key string) {

}
