package awsgo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecs"
)

var initErr error

// Client makes calls into AWS
type Client struct {
	config     *aws.Config
	session    *session.Session
	dryrunMode bool
	svc        *SVC
	awsContext *AWSContext
}

// NewClient creates a new Client
func NewClient(awsContext *AWSContext) (*Client, error) {
	var client Client
	creds, err := awsContext.Retrieve()
	if err != nil {
		return &client, err
	}
	awsConfig := aws.Config{
		Credentials: credentials.NewStaticCredentialsFromCreds(creds),
	}
	sess, err := session.NewSession(&awsConfig)
	if err != nil {
		return &client, err
	}
	client.awsContext = awsContext
	client.session = sess //sess.Copy()
	client.svc = &SVC{}
	return &client, nil
}

// AWSContext returns the client's AWSContext.
func (cl *Client) AWSContext() *AWSContext {
	return cl.awsContext
}

// AddConfig changes the underlying session with new Config options.
func (cl *Client) AddConfig(svcType ServiceType, options ConfigOptions) {
	cl.session = cl.session.Copy(options.GetDefaults(svcType))
}

// DryRunMode sets the DryRun bool
func (cl *Client) DryRunMode(enabled bool) {
	cl.dryrunMode = enabled
}

// EC2 returns the EC2 instance of the Client.
func (cl *Client) EC2() *ec2.EC2 {
	if cl.svc.ec2Svc == nil {
		cl.InitSVC(SvcTypeEC2)
	}
	return cl.svc.ec2Svc
}

// CW returns the CloudWatch instance of the Client.
func (cl *Client) CW() *cloudwatch.CloudWatch {
	if cl.svc.cwSvc == nil {
		cl.InitSVC(SvcTypeCloudWatch)
	}
	return cl.svc.cwSvc
}

// ECR returns the ECR instance of the Client.
func (cl *Client) ECR() *ecr.ECR {
	if cl.svc.ecrSvc == nil {
		cl.InitSVC(SvcTypeECR)
	}
	return cl.svc.ecrSvc
}

// ECS returns the ECS instance of the Client.
func (cl *Client) ECS() *ecs.ECS {
	if cl.svc.ecsSvc == nil {
		cl.InitSVC(SvcTypeECS)
	}
	return cl.svc.ecsSvc
}

// InitSVC inits the corresponding Service for the Client.
func (cl *Client) InitSVC(service ServiceType) {
	switch service {
	case SvcTypeEC2:
		cl.svc.ec2Svc = ec2.New(cl.session)
	case SvcTypeCloudWatch:
		cl.svc.cwSvc = cloudwatch.New(cl.session)
	case SvcTypeECR:
		cl.svc.ecrSvc = ecr.New(cl.session)
	case SvcTypeECS:
		cl.svc.ecsSvc = ecs.New(cl.session)
	}
}
