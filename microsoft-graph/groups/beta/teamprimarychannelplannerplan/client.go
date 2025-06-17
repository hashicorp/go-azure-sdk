package teamprimarychannelplannerplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelPlannerPlanClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelPlannerPlanClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelPlannerPlanClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelplannerplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelPlannerPlanClient: %+v", err)
	}

	return &TeamPrimaryChannelPlannerPlanClient{
		Client: client,
	}, nil
}
