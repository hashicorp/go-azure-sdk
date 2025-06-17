package teamchannelplannerplanbucket

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelPlannerPlanBucketClient struct {
	Client *msgraph.Client
}

func NewTeamChannelPlannerPlanBucketClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelPlannerPlanBucketClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelplannerplanbucket", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelPlannerPlanBucketClient: %+v", err)
	}

	return &TeamChannelPlannerPlanBucketClient{
		Client: client,
	}, nil
}
