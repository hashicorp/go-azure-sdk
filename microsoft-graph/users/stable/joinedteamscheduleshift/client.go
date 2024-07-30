package joinedteamscheduleshift

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamScheduleShiftClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamScheduleShiftClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamScheduleShiftClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamscheduleshift", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamScheduleShiftClient: %+v", err)
	}

	return &JoinedTeamScheduleShiftClient{
		Client: client,
	}, nil
}
