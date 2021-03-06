package awsgo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/sts"
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

// Session returns the base underlying session.
func (cl *Client) Session() *session.Session {
	return cl.session
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
		cl.InitSVC(SvcTypeCW)
	}
	return cl.svc.cwSvc
}

// CWLogs returns the CloudWatchLogs instance of the Client.
func (cl *Client) CWLogs() *cloudwatchlogs.CloudWatchLogs {
	if cl.svc.ecsSvc == nil {
		cl.InitSVC(SvcTypeCWLogs)
	}
	return cl.svc.cwlogsSvc
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

// STS returns the STS instance of the Client.
func (cl *Client) STS() *sts.STS {
	if cl.svc.stsSvc == nil {
		cl.InitSVC(SvcTypeSTS)
	}
	return cl.svc.stsSvc
}

// Kinesis returns the Kinesis instance of the Client.
func (cl *Client) Kinesis() *kinesis.Kinesis {
	if cl.svc.kinesisSvc == nil {
		cl.InitSVC(SvcTypeKinesis)
	}
	return cl.svc.kinesisSvc
}
