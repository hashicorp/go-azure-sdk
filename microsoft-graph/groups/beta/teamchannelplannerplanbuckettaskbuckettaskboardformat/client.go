package teamchannelplannerplanbuckettaskbuckettaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelPlannerPlanBucketTaskBucketTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewTeamChannelPlannerPlanBucketTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelPlannerPlanBucketTaskBucketTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelplannerplanbuckettaskbuckettaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelPlannerPlanBucketTaskBucketTaskBoardFormatClient: %+v", err)
	}

	return &TeamChannelPlannerPlanBucketTaskBucketTaskBoardFormatClient{
		Client: client,
	}, nil
}
