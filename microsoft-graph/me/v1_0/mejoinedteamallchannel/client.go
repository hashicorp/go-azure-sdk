package mejoinedteamallchannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamAllChannelClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamAllChannelClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamAllChannelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamallchannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamAllChannelClient: %+v", err)
	}

	return &MeJoinedTeamAllChannelClient{
		Client: client,
	}, nil
}
