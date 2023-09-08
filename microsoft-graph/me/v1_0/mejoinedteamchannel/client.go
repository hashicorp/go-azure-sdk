package mejoinedteamchannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamChannelClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamChannelClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamChannelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamchannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamChannelClient: %+v", err)
	}

	return &MeJoinedTeamChannelClient{
		Client: client,
	}, nil
}
