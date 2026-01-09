package teamchannelplannerplandetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelPlannerPlanDetailClient struct {
	Client *msgraph.Client
}

func NewTeamChannelPlannerPlanDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelPlannerPlanDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelplannerplandetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelPlannerPlanDetailClient: %+v", err)
	}

	return &TeamChannelPlannerPlanDetailClient{
		Client: client,
	}, nil
}
