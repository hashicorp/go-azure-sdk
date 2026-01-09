package joinedteamchanneltabteamsapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamChannelTabTeamsAppClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamChannelTabTeamsAppClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamChannelTabTeamsAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamchanneltabteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamChannelTabTeamsAppClient: %+v", err)
	}

	return &JoinedTeamChannelTabTeamsAppClient{
		Client: client,
	}, nil
}
