package userjoinedteamchannelsharedwithteamallowedmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamChannelSharedWithTeamAllowedMemberClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamChannelSharedWithTeamAllowedMemberClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamChannelSharedWithTeamAllowedMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamchannelsharedwithteamallowedmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamChannelSharedWithTeamAllowedMemberClient: %+v", err)
	}

	return &UserJoinedTeamChannelSharedWithTeamAllowedMemberClient{
		Client: client,
	}, nil
}
