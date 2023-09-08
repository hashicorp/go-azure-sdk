package mejoinedteamchannelsharedwithteamallowedmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamChannelSharedWithTeamAllowedMemberClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamChannelSharedWithTeamAllowedMemberClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamChannelSharedWithTeamAllowedMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamchannelsharedwithteamallowedmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamChannelSharedWithTeamAllowedMemberClient: %+v", err)
	}

	return &MeJoinedTeamChannelSharedWithTeamAllowedMemberClient{
		Client: client,
	}, nil
}
