package teamprimarychannelplanner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelPlannerClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelPlannerClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelPlannerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelplanner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelPlannerClient: %+v", err)
	}

	return &TeamPrimaryChannelPlannerClient{
		Client: client,
	}, nil
}
