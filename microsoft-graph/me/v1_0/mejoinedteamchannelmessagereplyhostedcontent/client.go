package mejoinedteamchannelmessagereplyhostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamChannelMessageReplyHostedContentClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamChannelMessageReplyHostedContentClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamChannelMessageReplyHostedContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamchannelmessagereplyhostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamChannelMessageReplyHostedContentClient: %+v", err)
	}

	return &MeJoinedTeamChannelMessageReplyHostedContentClient{
		Client: client,
	}, nil
}
