package userjoinedteamscheduletimesoff

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamScheduleTimesOffClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamScheduleTimesOffClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamScheduleTimesOffClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamscheduletimesoff", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamScheduleTimesOffClient: %+v", err)
	}

	return &UserJoinedTeamScheduleTimesOffClient{
		Client: client,
	}, nil
}
