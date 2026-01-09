package joinedteamchannelsharedwithteamteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamChannelSharedWithTeamTeamClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamChannelSharedWithTeamTeamClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamChannelSharedWithTeamTeamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamchannelsharedwithteamteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamChannelSharedWithTeamTeamClient: %+v", err)
	}

	return &JoinedTeamChannelSharedWithTeamTeamClient{
		Client: client,
	}, nil
}
