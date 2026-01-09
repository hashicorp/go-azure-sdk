package joinedteamscheduletimecard

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamScheduleTimeCardClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamScheduleTimeCardClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamScheduleTimeCardClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamscheduletimecard", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamScheduleTimeCardClient: %+v", err)
	}

	return &JoinedTeamScheduleTimeCardClient{
		Client: client,
	}, nil
}
