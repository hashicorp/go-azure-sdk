package userjoinedteamchannelmessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamChannelMessageClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamChannelMessageClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamChannelMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamchannelmessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamChannelMessageClient: %+v", err)
	}

	return &UserJoinedTeamChannelMessageClient{
		Client: client,
	}, nil
}
