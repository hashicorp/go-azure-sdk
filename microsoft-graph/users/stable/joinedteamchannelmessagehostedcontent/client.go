package joinedteamchannelmessagehostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamChannelMessageHostedContentClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamChannelMessageHostedContentClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamChannelMessageHostedContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamchannelmessagehostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamChannelMessageHostedContentClient: %+v", err)
	}

	return &JoinedTeamChannelMessageHostedContentClient{
		Client: client,
	}, nil
}
