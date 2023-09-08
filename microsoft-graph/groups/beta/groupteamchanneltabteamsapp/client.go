package groupteamchanneltabteamsapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamChannelTabTeamsAppClient struct {
	Client *msgraph.Client
}

func NewGroupTeamChannelTabTeamsAppClientWithBaseURI(api sdkEnv.Api) (*GroupTeamChannelTabTeamsAppClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamchanneltabteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamChannelTabTeamsAppClient: %+v", err)
	}

	return &GroupTeamChannelTabTeamsAppClient{
		Client: client,
	}, nil
}
