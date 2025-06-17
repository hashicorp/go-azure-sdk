package teamprimarychannelplannerplantask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelPlannerPlanTaskClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelPlannerPlanTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelPlannerPlanTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelplannerplantask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelPlannerPlanTaskClient: %+v", err)
	}

	return &TeamPrimaryChannelPlannerPlanTaskClient{
		Client: client,
	}, nil
}
