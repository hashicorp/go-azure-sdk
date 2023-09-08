package mejoinedteamchannelmessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamChannelMessageClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamChannelMessageClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamChannelMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamchannelmessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamChannelMessageClient: %+v", err)
	}

	return &MeJoinedTeamChannelMessageClient{
		Client: client,
	}, nil
}
