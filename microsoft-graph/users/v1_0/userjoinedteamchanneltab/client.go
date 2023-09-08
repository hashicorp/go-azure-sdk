package userjoinedteamchanneltab

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamChannelTabClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamChannelTabClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamChannelTabClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamchanneltab", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamChannelTabClient: %+v", err)
	}

	return &UserJoinedTeamChannelTabClient{
		Client: client,
	}, nil
}
