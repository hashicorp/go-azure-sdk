package teamprimarychannelplannerplanbuckettaskdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelPlannerPlanBucketTaskDetailClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelPlannerPlanBucketTaskDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelPlannerPlanBucketTaskDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelplannerplanbuckettaskdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelPlannerPlanBucketTaskDetailClient: %+v", err)
	}

	return &TeamPrimaryChannelPlannerPlanBucketTaskDetailClient{
		Client: client,
	}, nil
}
