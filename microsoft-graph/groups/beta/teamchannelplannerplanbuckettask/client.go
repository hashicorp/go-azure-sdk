package teamchannelplannerplanbuckettask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelPlannerPlanBucketTaskClient struct {
	Client *msgraph.Client
}

func NewTeamChannelPlannerPlanBucketTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelPlannerPlanBucketTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelplannerplanbuckettask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelPlannerPlanBucketTaskClient: %+v", err)
	}

	return &TeamChannelPlannerPlanBucketTaskClient{
		Client: client,
	}, nil
}
