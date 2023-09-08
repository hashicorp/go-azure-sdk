package userjoinedteamprimarychannelsharedwithteamallowedmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamprimarychannelsharedwithteamallowedmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient: %+v", err)
	}

	return &UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient{
		Client: client,
	}, nil
}
