package teamchannelplannerplantaskdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelPlannerPlanTaskDetailClient struct {
	Client *msgraph.Client
}

func NewTeamChannelPlannerPlanTaskDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelPlannerPlanTaskDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelplannerplantaskdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelPlannerPlanTaskDetailClient: %+v", err)
	}

	return &TeamChannelPlannerPlanTaskDetailClient{
		Client: client,
	}, nil
}
