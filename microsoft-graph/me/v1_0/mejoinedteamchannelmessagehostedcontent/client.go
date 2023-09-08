package mejoinedteamchannelmessagehostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamChannelMessageHostedContentClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamChannelMessageHostedContentClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamChannelMessageHostedContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamchannelmessagehostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamChannelMessageHostedContentClient: %+v", err)
	}

	return &MeJoinedTeamChannelMessageHostedContentClient{
		Client: client,
	}, nil
}
