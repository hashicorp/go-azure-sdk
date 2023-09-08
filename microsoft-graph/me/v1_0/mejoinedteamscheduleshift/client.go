package mejoinedteamscheduleshift

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamScheduleShiftClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamScheduleShiftClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamScheduleShiftClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamscheduleshift", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamScheduleShiftClient: %+v", err)
	}

	return &MeJoinedTeamScheduleShiftClient{
		Client: client,
	}, nil
}
