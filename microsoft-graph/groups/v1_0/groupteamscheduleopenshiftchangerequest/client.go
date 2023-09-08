package groupteamscheduleopenshiftchangerequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamScheduleOpenShiftChangeRequestClient struct {
	Client *msgraph.Client
}

func NewGroupTeamScheduleOpenShiftChangeRequestClientWithBaseURI(api sdkEnv.Api) (*GroupTeamScheduleOpenShiftChangeRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamscheduleopenshiftchangerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamScheduleOpenShiftChangeRequestClient: %+v", err)
	}

	return &GroupTeamScheduleOpenShiftChangeRequestClient{
		Client: client,
	}, nil
}
