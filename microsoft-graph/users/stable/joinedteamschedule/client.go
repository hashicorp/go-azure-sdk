package joinedteamschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamScheduleClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamScheduleClient: %+v", err)
	}

	return &JoinedTeamScheduleClient{
		Client: client,
	}, nil
}
