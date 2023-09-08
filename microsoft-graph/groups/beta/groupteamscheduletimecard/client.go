package groupteamscheduletimecard

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamScheduleTimeCardClient struct {
	Client *msgraph.Client
}

func NewGroupTeamScheduleTimeCardClientWithBaseURI(api sdkEnv.Api) (*GroupTeamScheduleTimeCardClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamscheduletimecard", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamScheduleTimeCardClient: %+v", err)
	}

	return &GroupTeamScheduleTimeCardClient{
		Client: client,
	}, nil
}
