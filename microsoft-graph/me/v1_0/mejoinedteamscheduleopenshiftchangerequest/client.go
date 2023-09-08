package mejoinedteamscheduleopenshiftchangerequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamScheduleOpenShiftChangeRequestClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamScheduleOpenShiftChangeRequestClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamScheduleOpenShiftChangeRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamscheduleopenshiftchangerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamScheduleOpenShiftChangeRequestClient: %+v", err)
	}

	return &MeJoinedTeamScheduleOpenShiftChangeRequestClient{
		Client: client,
	}, nil
}
