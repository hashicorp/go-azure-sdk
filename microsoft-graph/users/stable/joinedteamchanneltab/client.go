package joinedteamchanneltab

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamChannelTabClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamChannelTabClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamChannelTabClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamchanneltab", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamChannelTabClient: %+v", err)
	}

	return &JoinedTeamChannelTabClient{
		Client: client,
	}, nil
}
