package teamprimarychannelplannerplantaskdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelPlannerPlanTaskDetailClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelPlannerPlanTaskDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelPlannerPlanTaskDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelplannerplantaskdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelPlannerPlanTaskDetailClient: %+v", err)
	}

	return &TeamPrimaryChannelPlannerPlanTaskDetailClient{
		Client: client,
	}, nil
}
