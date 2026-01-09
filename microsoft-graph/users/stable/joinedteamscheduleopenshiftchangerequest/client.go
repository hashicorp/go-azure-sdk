package joinedteamscheduleopenshiftchangerequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamScheduleOpenShiftChangeRequestClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamScheduleOpenShiftChangeRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamScheduleOpenShiftChangeRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamscheduleopenshiftchangerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamScheduleOpenShiftChangeRequestClient: %+v", err)
	}

	return &JoinedTeamScheduleOpenShiftChangeRequestClient{
		Client: client,
	}, nil
}
