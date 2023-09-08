package mejoinedteamscheduletimeoffreason

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamScheduleTimeOffReasonClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamScheduleTimeOffReasonClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamScheduleTimeOffReasonClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamscheduletimeoffreason", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamScheduleTimeOffReasonClient: %+v", err)
	}

	return &MeJoinedTeamScheduleTimeOffReasonClient{
		Client: client,
	}, nil
}
