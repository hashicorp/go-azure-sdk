package mejoinedteamprimarychannelsharedwithteamallowedmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamprimarychannelsharedwithteamallowedmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient: %+v", err)
	}

	return &MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient{
		Client: client,
	}, nil
}
