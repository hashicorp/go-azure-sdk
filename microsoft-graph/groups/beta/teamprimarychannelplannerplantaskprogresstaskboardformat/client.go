package teamprimarychannelplannerplantaskprogresstaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelPlannerPlanTaskProgressTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelPlannerPlanTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelPlannerPlanTaskProgressTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelplannerplantaskprogresstaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelPlannerPlanTaskProgressTaskBoardFormatClient: %+v", err)
	}

	return &TeamPrimaryChannelPlannerPlanTaskProgressTaskBoardFormatClient{
		Client: client,
	}, nil
}
