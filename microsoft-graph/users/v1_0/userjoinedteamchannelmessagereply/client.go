package userjoinedteamchannelmessagereply

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamChannelMessageReplyClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamChannelMessageReplyClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamChannelMessageReplyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamchannelmessagereply", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamChannelMessageReplyClient: %+v", err)
	}

	return &UserJoinedTeamChannelMessageReplyClient{
		Client: client,
	}, nil
}
