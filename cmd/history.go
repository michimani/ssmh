package cmd

import (
	"context"
	"errors"
	"ssmhistory/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	ssmtypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

var initialToken = "init"

func ListSessionHistory(ctx context.Context, c *types.SSMHistoryClient, p *types.HistoryParams) error {
	list, err := describeSessions(ctx, c, p.MaxResults)
	if err != nil {
		return err
	}

	list.Display()

	return nil
}

func describeSessions(ctx context.Context, c *types.SSMHistoryClient, maxResults *types.MaxResults) (types.SSMHistoryItemList, error) {
	list := types.SSMHistoryItemList{}

	nextToken := &initialToken
	for nextToken != nil {
		in := &ssm.DescribeSessionsInput{
			State: ssmtypes.SessionStateHistory,
		}
		if maxResults.Valid() {
			in.MaxResults = int32(*maxResults)
		}

		out, err := c.SSM.DescribeSessions(ctx, in)
		if err != nil {
			return nil, err
		}

		for _, h := range out.Sessions {
			instanceName, err := getEC2InstanceName(ctx, c.EC2, h.Target)
			if err != nil {
				return nil, err
			}

			item := types.SSMHistoryItem{
				Target:       aws.ToString(h.Target),
				InstanceName: instanceName,
				SessionID:    aws.ToString(h.SessionId),
				Reason:       aws.ToString(h.Reason),
				StartDate:    h.StartDate,
				EndDate:      h.EndDate,
			}

			list = append(list, item)
		}

		nextToken = out.NextToken
		if maxResults.Valid() {
			nextToken = nil
		}
	}

	return list, nil
}

func getEC2InstanceName(ctx context.Context, c *ec2.Client, instanceID *string) (string, error) {
	if instanceID == nil {
		return "", errors.New("instanceID is nil")
	}

	in := &ec2.DescribeInstancesInput{
		InstanceIds: []string{aws.ToString(instanceID)},
	}

	out, err := c.DescribeInstances(ctx, in)
	if err != nil {
		return "", err
	}

	if len(out.Reservations) == 0 {
		return "", errors.New("no Reservations")
	}

	if len(out.Reservations[0].Instances) == 0 {
		return "", errors.New("no Instances")
	}

	instance := out.Reservations[0].Instances[0]

	if len(instance.Tags) == 0 {
		// no tags
		return "", nil
	}

	name := ""
	for _, t := range instance.Tags {
		if *t.Key == "Name" {
			name = *t.Value
			break
		}
	}

	return name, nil
}
