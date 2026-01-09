package teamprimarychannelplannerplantaskassignedtotaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelPlannerPlanTaskAssignedToTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelPlannerPlanTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelPlannerPlanTaskAssignedToTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelplannerplantaskassignedtotaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelPlannerPlanTaskAssignedToTaskBoardFormatClient: %+v", err)
	}

	return &TeamPrimaryChannelPlannerPlanTaskAssignedToTaskBoardFormatClient{
		Client: client,
	}, nil
}
