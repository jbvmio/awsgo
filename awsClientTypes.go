package awsgo

import (
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecs"
)

// ServiceType defines a specific AWS Service available for a Client.
type ServiceType int

// DefaultConfig returns the default config file name for a ServiceType.
func (s ServiceType) DefaultConfig() string {
	return svcTypeConfigName[s]
}

// Available Service Types:
const (
	SvcTypeEC2 ServiceType = iota // 0
	SvcTypeCW
	SvcTypeCWLogs
	SvcTypeECR
	SvcTypeECS
)

var svcTypeConfigName = [...]string{
	"ec2_defaults",
	"cw_defaults",
	"cwlogs_defaults",
	"ecr_defaults",
	"ecs_defaults",
}

// SVC contains available AWS service clients
type SVC struct {
	ec2Svc    *ec2.EC2
	cwSvc     *cloudwatch.CloudWatch
	cwlogsSvc *cloudwatchlogs.CloudWatchLogs
	ecrSvc    *ecr.ECR
	ecsSvc    *ecs.ECS
}

// InitSVC inits the corresponding Service for the Client.
func (cl *Client) InitSVC(service ServiceType) {
	switch service {
	case SvcTypeEC2:
		cl.svc.ec2Svc = ec2.New(cl.session)
	case SvcTypeCW:
		cl.svc.cwSvc = cloudwatch.New(cl.session)
	case SvcTypeCWLogs:
		cl.svc.cwlogsSvc = cloudwatchlogs.New(cl.session)
	case SvcTypeECR:
		cl.svc.ecrSvc = ecr.New(cl.session)
	case SvcTypeECS:
		cl.svc.ecsSvc = ecs.New(cl.session)
	}
}
