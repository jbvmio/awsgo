package awsgo

import "github.com/aws/aws-sdk-go/aws"

// DefaultConfigName holds all defaults if there is not any overrides specifid for a ServiceType.
const DefaultConfigName = `defaults`

// ConfigOptions contains the options that deliver configuration details.
type ConfigOptions interface {
	GetDefaults(ServiceType) *aws.Config
}

// GetDefault returns the default value for the given key (TODO).
func GetDefault(dir, key string) {

}
