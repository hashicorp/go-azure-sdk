package userjoinedteamscheduleopenshiftchangerequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamScheduleOpenShiftChangeRequestClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamScheduleOpenShiftChangeRequestClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamScheduleOpenShiftChangeRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamscheduleopenshiftchangerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamScheduleOpenShiftChangeRequestClient: %+v", err)
	}

	return &UserJoinedTeamScheduleOpenShiftChangeRequestClient{
		Client: client,
	}, nil
}
