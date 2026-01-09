package teamchannelsharedwithteamallowedmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelSharedWithTeamAllowedMemberClient struct {
	Client *msgraph.Client
}

func NewTeamChannelSharedWithTeamAllowedMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelSharedWithTeamAllowedMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelsharedwithteamallowedmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelSharedWithTeamAllowedMemberClient: %+v", err)
	}

	return &TeamChannelSharedWithTeamAllowedMemberClient{
		Client: client,
	}, nil
}
