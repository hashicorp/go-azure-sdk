package teamprimarychannelplannerplandetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelPlannerPlanDetailClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelPlannerPlanDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelPlannerPlanDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelplannerplandetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelPlannerPlanDetailClient: %+v", err)
	}

	return &TeamPrimaryChannelPlannerPlanDetailClient{
		Client: client,
	}, nil
}
