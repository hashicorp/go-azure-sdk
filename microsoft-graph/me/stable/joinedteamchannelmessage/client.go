package joinedteamchannelmessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamChannelMessageClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamChannelMessageClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamChannelMessageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamchannelmessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamChannelMessageClient: %+v", err)
	}

	return &JoinedTeamChannelMessageClient{
		Client: client,
	}, nil
}
