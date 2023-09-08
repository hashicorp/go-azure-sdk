package userjoinedteamscheduleshift

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamScheduleShiftClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamScheduleShiftClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamScheduleShiftClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamscheduleshift", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamScheduleShiftClient: %+v", err)
	}

	return &UserJoinedTeamScheduleShiftClient{
		Client: client,
	}, nil
}
