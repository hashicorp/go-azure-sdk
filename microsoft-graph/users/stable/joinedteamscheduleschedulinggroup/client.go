package joinedteamscheduleschedulinggroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamScheduleSchedulingGroupClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamScheduleSchedulingGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamScheduleSchedulingGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamscheduleschedulinggroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamScheduleSchedulingGroupClient: %+v", err)
	}

	return &JoinedTeamScheduleSchedulingGroupClient{
		Client: client,
	}, nil
}
