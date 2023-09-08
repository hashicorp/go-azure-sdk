package groupteamscheduletimeoffrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamScheduleTimeOffRequestClient struct {
	Client *msgraph.Client
}

func NewGroupTeamScheduleTimeOffRequestClientWithBaseURI(api sdkEnv.Api) (*GroupTeamScheduleTimeOffRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamscheduletimeoffrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamScheduleTimeOffRequestClient: %+v", err)
	}

	return &GroupTeamScheduleTimeOffRequestClient{
		Client: client,
	}, nil
}
