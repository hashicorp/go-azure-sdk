package userjoinedteamincomingchannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamIncomingChannelClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamIncomingChannelClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamIncomingChannelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamincomingchannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamIncomingChannelClient: %+v", err)
	}

	return &UserJoinedTeamIncomingChannelClient{
		Client: client,
	}, nil
}
