package mejoinedteamchanneltab

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamChannelTabClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamChannelTabClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamChannelTabClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamchanneltab", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamChannelTabClient: %+v", err)
	}

	return &MeJoinedTeamChannelTabClient{
		Client: client,
	}, nil
}
