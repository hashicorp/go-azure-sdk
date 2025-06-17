package teamchannelplannerplanbuckettaskdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelPlannerPlanBucketTaskDetailClient struct {
	Client *msgraph.Client
}

func NewTeamChannelPlannerPlanBucketTaskDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelPlannerPlanBucketTaskDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelplannerplanbuckettaskdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelPlannerPlanBucketTaskDetailClient: %+v", err)
	}

	return &TeamChannelPlannerPlanBucketTaskDetailClient{
		Client: client,
	}, nil
}
