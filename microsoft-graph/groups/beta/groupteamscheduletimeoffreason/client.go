package groupteamscheduletimeoffreason

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamScheduleTimeOffReasonClient struct {
	Client *msgraph.Client
}

func NewGroupTeamScheduleTimeOffReasonClientWithBaseURI(api sdkEnv.Api) (*GroupTeamScheduleTimeOffReasonClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamscheduletimeoffreason", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamScheduleTimeOffReasonClient: %+v", err)
	}

	return &GroupTeamScheduleTimeOffReasonClient{
		Client: client,
	}, nil
}
