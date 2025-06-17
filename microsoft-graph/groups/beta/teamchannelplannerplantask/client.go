package teamchannelplannerplantask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelPlannerPlanTaskClient struct {
	Client *msgraph.Client
}

func NewTeamChannelPlannerPlanTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelPlannerPlanTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelplannerplantask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelPlannerPlanTaskClient: %+v", err)
	}

	return &TeamChannelPlannerPlanTaskClient{
		Client: client,
	}, nil
}
