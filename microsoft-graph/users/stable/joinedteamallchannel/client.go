package joinedteamallchannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamAllChannelClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamAllChannelClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamAllChannelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamallchannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamAllChannelClient: %+v", err)
	}

	return &JoinedTeamAllChannelClient{
		Client: client,
	}, nil
}
