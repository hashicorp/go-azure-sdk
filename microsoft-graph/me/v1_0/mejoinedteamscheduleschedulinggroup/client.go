package mejoinedteamscheduleschedulinggroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamScheduleSchedulingGroupClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamScheduleSchedulingGroupClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamScheduleSchedulingGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamscheduleschedulinggroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamScheduleSchedulingGroupClient: %+v", err)
	}

	return &MeJoinedTeamScheduleSchedulingGroupClient{
		Client: client,
	}, nil
}
