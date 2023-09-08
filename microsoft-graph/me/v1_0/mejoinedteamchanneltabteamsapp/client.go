package mejoinedteamchanneltabteamsapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamChannelTabTeamsAppClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamChannelTabTeamsAppClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamChannelTabTeamsAppClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamchanneltabteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamChannelTabTeamsAppClient: %+v", err)
	}

	return &MeJoinedTeamChannelTabTeamsAppClient{
		Client: client,
	}, nil
}
