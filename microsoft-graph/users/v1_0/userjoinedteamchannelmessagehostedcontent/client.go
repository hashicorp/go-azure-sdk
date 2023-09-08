package userjoinedteamchannelmessagehostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamChannelMessageHostedContentClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamChannelMessageHostedContentClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamChannelMessageHostedContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamchannelmessagehostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamChannelMessageHostedContentClient: %+v", err)
	}

	return &UserJoinedTeamChannelMessageHostedContentClient{
		Client: client,
	}, nil
}
