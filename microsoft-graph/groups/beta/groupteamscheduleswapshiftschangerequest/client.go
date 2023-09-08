package groupteamscheduleswapshiftschangerequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamScheduleSwapShiftsChangeRequestClient struct {
	Client *msgraph.Client
}

func NewGroupTeamScheduleSwapShiftsChangeRequestClientWithBaseURI(api sdkEnv.Api) (*GroupTeamScheduleSwapShiftsChangeRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamscheduleswapshiftschangerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamScheduleSwapShiftsChangeRequestClient: %+v", err)
	}

	return &GroupTeamScheduleSwapShiftsChangeRequestClient{
		Client: client,
	}, nil
}
