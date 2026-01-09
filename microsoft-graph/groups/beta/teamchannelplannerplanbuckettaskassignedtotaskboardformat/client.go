package teamchannelplannerplanbuckettaskassignedtotaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelPlannerPlanBucketTaskAssignedToTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewTeamChannelPlannerPlanBucketTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelPlannerPlanBucketTaskAssignedToTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelplannerplanbuckettaskassignedtotaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelPlannerPlanBucketTaskAssignedToTaskBoardFormatClient: %+v", err)
	}

	return &TeamChannelPlannerPlanBucketTaskAssignedToTaskBoardFormatClient{
		Client: client,
	}, nil
}
