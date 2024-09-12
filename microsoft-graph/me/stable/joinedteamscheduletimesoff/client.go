package joinedteamscheduletimesoff

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamScheduleTimesOffClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamScheduleTimesOffClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamScheduleTimesOffClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamscheduletimesoff", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamScheduleTimesOffClient: %+v", err)
	}

	return &JoinedTeamScheduleTimesOffClient{
		Client: client,
	}, nil
}
