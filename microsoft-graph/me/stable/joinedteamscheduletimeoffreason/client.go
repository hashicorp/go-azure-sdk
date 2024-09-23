package joinedteamscheduletimeoffreason

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamScheduleTimeOffReasonClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamScheduleTimeOffReasonClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamScheduleTimeOffReasonClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamscheduletimeoffreason", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamScheduleTimeOffReasonClient: %+v", err)
	}

	return &JoinedTeamScheduleTimeOffReasonClient{
		Client: client,
	}, nil
}
