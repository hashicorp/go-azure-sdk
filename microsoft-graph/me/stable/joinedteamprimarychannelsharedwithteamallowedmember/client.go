package joinedteamprimarychannelsharedwithteamallowedmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamprimarychannelsharedwithteamallowedmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient: %+v", err)
	}

	return &JoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient{
		Client: client,
	}, nil
}
