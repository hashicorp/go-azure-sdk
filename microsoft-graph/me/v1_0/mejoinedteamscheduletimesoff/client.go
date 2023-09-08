package mejoinedteamscheduletimesoff

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamScheduleTimesOffClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamScheduleTimesOffClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamScheduleTimesOffClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamscheduletimesoff", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamScheduleTimesOffClient: %+v", err)
	}

	return &MeJoinedTeamScheduleTimesOffClient{
		Client: client,
	}, nil
}
