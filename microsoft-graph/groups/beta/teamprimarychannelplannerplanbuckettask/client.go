package teamprimarychannelplannerplanbuckettask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelPlannerPlanBucketTaskClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelPlannerPlanBucketTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelPlannerPlanBucketTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelplannerplanbuckettask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelPlannerPlanBucketTaskClient: %+v", err)
	}

	return &TeamPrimaryChannelPlannerPlanBucketTaskClient{
		Client: client,
	}, nil
}
