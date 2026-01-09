package joinedteamchannelsharedwithteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamChannelSharedWithTeamClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamChannelSharedWithTeamClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamChannelSharedWithTeamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamchannelsharedwithteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamChannelSharedWithTeamClient: %+v", err)
	}

	return &JoinedTeamChannelSharedWithTeamClient{
		Client: client,
	}, nil
}
