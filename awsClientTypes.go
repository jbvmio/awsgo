package awsgo

import (
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecr"
)

// ServiceType defines a specific AWS Service available for a Client.
type ServiceType int

// DefaultConfig returns the default config file name for a ServiceType.
func (s ServiceType) DefaultConfig() string {
	return svcTypeConfigName[s]
}

// Available Service Types:
const (
	SvcTypeEC2        ServiceType = 0
	SvcTypeCloudWatch ServiceType = 1
	SvcTypeECR        ServiceType = 2
)

var svcTypeConfigName = [...]string{
	"ec2_defaults",
	"cw_defaults",
	"ecr_defaults",
}

// SVC contains available AWS service clients
type SVC struct {
	ec2Svc *ec2.EC2
	cwSvc  *cloudwatch.CloudWatch
	ecrSvc *ecr.ECR
}
