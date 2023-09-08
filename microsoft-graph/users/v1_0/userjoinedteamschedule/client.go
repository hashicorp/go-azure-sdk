package userjoinedteamschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamScheduleClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamScheduleClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamScheduleClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamScheduleClient: %+v", err)
	}

	return &UserJoinedTeamScheduleClient{
		Client: client,
	}, nil
}
