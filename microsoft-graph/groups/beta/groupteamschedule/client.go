package groupteamschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamScheduleClient struct {
	Client *msgraph.Client
}

func NewGroupTeamScheduleClientWithBaseURI(api sdkEnv.Api) (*GroupTeamScheduleClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamScheduleClient: %+v", err)
	}

	return &GroupTeamScheduleClient{
		Client: client,
	}, nil
}
