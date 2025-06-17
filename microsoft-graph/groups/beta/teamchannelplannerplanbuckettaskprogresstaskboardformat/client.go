package teamchannelplannerplanbuckettaskprogresstaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelPlannerPlanBucketTaskProgressTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewTeamChannelPlannerPlanBucketTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelPlannerPlanBucketTaskProgressTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelplannerplanbuckettaskprogresstaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelPlannerPlanBucketTaskProgressTaskBoardFormatClient: %+v", err)
	}

	return &TeamChannelPlannerPlanBucketTaskProgressTaskBoardFormatClient{
		Client: client,
	}, nil
}
