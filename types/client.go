package types

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type SSMHistoryClient struct {
	SSM *ssm.Client
	EC2 *ec2.Client
}

type NewClientInput struct {
	Context context.Context
	Region  string
}
