package userjoinedteamscheduleopenshift

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamScheduleOpenShiftClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamScheduleOpenShiftClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamScheduleOpenShiftClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamscheduleopenshift", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamScheduleOpenShiftClient: %+v", err)
	}

	return &UserJoinedTeamScheduleOpenShiftClient{
		Client: client,
	}, nil
}
