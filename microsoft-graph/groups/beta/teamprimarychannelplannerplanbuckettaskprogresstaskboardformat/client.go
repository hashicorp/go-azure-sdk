package teamprimarychannelplannerplanbuckettaskprogresstaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelplannerplanbuckettaskprogresstaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatClient: %+v", err)
	}

	return &TeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatClient{
		Client: client,
	}, nil
}
