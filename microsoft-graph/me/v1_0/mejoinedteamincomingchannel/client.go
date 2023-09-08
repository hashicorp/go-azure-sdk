package mejoinedteamincomingchannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamIncomingChannelClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamIncomingChannelClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamIncomingChannelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamincomingchannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamIncomingChannelClient: %+v", err)
	}

	return &MeJoinedTeamIncomingChannelClient{
		Client: client,
	}, nil
}
