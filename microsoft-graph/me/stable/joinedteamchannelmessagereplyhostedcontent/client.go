package joinedteamchannelmessagereplyhostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamChannelMessageReplyHostedContentClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamChannelMessageReplyHostedContentClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamChannelMessageReplyHostedContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamchannelmessagereplyhostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamChannelMessageReplyHostedContentClient: %+v", err)
	}

	return &JoinedTeamChannelMessageReplyHostedContentClient{
		Client: client,
	}, nil
}
