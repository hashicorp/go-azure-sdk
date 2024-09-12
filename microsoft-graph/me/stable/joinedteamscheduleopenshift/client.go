package joinedteamscheduleopenshift

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamScheduleOpenShiftClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamScheduleOpenShiftClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamScheduleOpenShiftClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamscheduleopenshift", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamScheduleOpenShiftClient: %+v", err)
	}

	return &JoinedTeamScheduleOpenShiftClient{
		Client: client,
	}, nil
}
