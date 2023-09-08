package mejoinedteamscheduletimeoffrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamScheduleTimeOffRequestClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamScheduleTimeOffRequestClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamScheduleTimeOffRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamscheduletimeoffrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamScheduleTimeOffRequestClient: %+v", err)
	}

	return &MeJoinedTeamScheduleTimeOffRequestClient{
		Client: client,
	}, nil
}
