package internal

import (
	"ssmhistory/types"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func NewClient(in *types.NewClientInput) (*types.SSMHistoryClient, error) {
	sc, err := newSSMClient(in)
	if err != nil {
		return nil, err
	}

	ec, err := newEC2Client(in)
	if err != nil {
		return nil, err
	}

	c := &types.SSMHistoryClient{
		SSM: sc,
		EC2: ec,
	}

	return c, nil
}

func newEC2Client(in *types.NewClientInput) (*ec2.Client, error) {
	cfg, err := config.LoadDefaultConfig(in.Context,
		config.WithRegion(in.Region),
	)
	if err != nil {
		return nil, err
	}

	c := ec2.NewFromConfig(cfg)
	return c, nil
}

func newSSMClient(in *types.NewClientInput) (*ssm.Client, error) {
	cfg, err := config.LoadDefaultConfig(in.Context,
		config.WithRegion(in.Region),
	)
	if err != nil {
		return nil, err
	}

	c := ssm.NewFromConfig(cfg)
	return c, nil
}
