package mejoinedteamchannelmessagereply

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamChannelMessageReplyClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamChannelMessageReplyClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamChannelMessageReplyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamchannelmessagereply", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamChannelMessageReplyClient: %+v", err)
	}

	return &MeJoinedTeamChannelMessageReplyClient{
		Client: client,
	}, nil
}
