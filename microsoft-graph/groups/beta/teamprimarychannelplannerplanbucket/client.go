package teamprimarychannelplannerplanbucket

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelPlannerPlanBucketClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelPlannerPlanBucketClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelPlannerPlanBucketClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelplannerplanbucket", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelPlannerPlanBucketClient: %+v", err)
	}

	return &TeamPrimaryChannelPlannerPlanBucketClient{
		Client: client,
	}, nil
}
