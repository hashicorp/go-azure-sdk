package teamprimarychannelplannerplanbuckettaskassignedtotaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelPlannerPlanBucketTaskAssignedToTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelPlannerPlanBucketTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelPlannerPlanBucketTaskAssignedToTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelplannerplanbuckettaskassignedtotaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelPlannerPlanBucketTaskAssignedToTaskBoardFormatClient: %+v", err)
	}

	return &TeamPrimaryChannelPlannerPlanBucketTaskAssignedToTaskBoardFormatClient{
		Client: client,
	}, nil
}
