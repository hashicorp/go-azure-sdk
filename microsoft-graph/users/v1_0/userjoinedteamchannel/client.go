package userjoinedteamchannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamChannelClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamChannelClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamChannelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamchannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamChannelClient: %+v", err)
	}

	return &UserJoinedTeamChannelClient{
		Client: client,
	}, nil
}
