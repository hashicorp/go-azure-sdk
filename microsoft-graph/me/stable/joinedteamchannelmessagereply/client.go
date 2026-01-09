package joinedteamchannelmessagereply

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamChannelMessageReplyClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamChannelMessageReplyClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamChannelMessageReplyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamchannelmessagereply", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamChannelMessageReplyClient: %+v", err)
	}

	return &JoinedTeamChannelMessageReplyClient{
		Client: client,
	}, nil
}
