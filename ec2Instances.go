package awsgo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// GetInstances returns Instances based on the given ids or all if no ids are given.
func (cl *Client) GetInstances(ids ...string) []Instance {
	return cl.GetInstanceMap().GetInstances(ids...)
}

// GetInstances returns Instances based on the given ids or all if no ids are given.
func (i InstanceMap) GetInstances(ids ...string) []Instance {
	var instances []Instance
	switch {
	case len(ids) > 0:
		for _, id := range ids {
			if i[id] != nil {
				var inst Instance
				inst.convertFrom(i[id])
				instances = append(instances, inst)
			}
		}
	default:
		for id := range i {
			var inst Instance
			inst.convertFrom(i[id])
			instances = append(instances, inst)
		}
	}
	return instances
}

// StartEC2Instances starts AWS Instances by id.
func (cl *Client) StartEC2Instances(ids ...string) []InstanceStateChange {
	var stateChanges []InstanceStateChange //*ec2.StartInstancesOutput
	var instances []*string
	for _, i := range ids {
		instances = append(instances, &i)
	}
	input := &ec2.StartInstancesInput{
		InstanceIds: instances,
		DryRun:      aws.Bool(cl.dryrunMode),
	}
	output, err := cl.EC2().StartInstances(input)
	if aerr, ok := err.(awserr.Error); ok {
		fmt.Printf("%v\n%v\n%v\n", aerr.Code(), aerr.Message(), err)
	}
	if err == nil {
		for _, out := range output.StartingInstances {
			var isc InstanceStateChange
			isc.convertFrom(out)
			stateChanges = append(stateChanges, isc)
		}
	}
	return stateChanges
}

// StopEC2Instances stops AWS Instances by id.
func (cl *Client) StopEC2Instances(ids ...string) []InstanceStateChange {
	var stateChanges []InstanceStateChange //*ec2.StartInstancesOutput
	var instances []*string
	for _, i := range ids {
		instances = append(instances, &i)
	}
	input := &ec2.StopInstancesInput{
		InstanceIds: instances,
		DryRun:      aws.Bool(cl.dryrunMode),
	}
	output, err := cl.EC2().StopInstances(input)
	if aerr, ok := err.(awserr.Error); ok {
		fmt.Printf("%v\n%v\n%v\n", aerr.Code(), aerr.Message(), err)
	}
	if err == nil {
		for _, out := range output.StoppingInstances {
			var isc InstanceStateChange
			isc.convertFrom(out)
			stateChanges = append(stateChanges, isc)
		}
	}
	return stateChanges
}

// RebootEC2Instances reboots AWS Instances by id, returns true if request successfully sent, false if not.
func (cl *Client) RebootEC2Instances(ids ...string) bool {
	//var stateChanges []InstanceStateChange //*ec2.StartInstancesOutput
	var instances []*string
	for _, i := range ids {
		instances = append(instances, &i)
	}
	input := &ec2.RebootInstancesInput{
		InstanceIds: instances,
		DryRun:      aws.Bool(cl.dryrunMode),
	}
	_, err := cl.EC2().RebootInstances(input)
	if aerr, ok := err.(awserr.Error); ok {
		fmt.Printf("%v\n%v\n%v\n", aerr.Code(), aerr.Message(), err)
		return false
	}
	return true
}
