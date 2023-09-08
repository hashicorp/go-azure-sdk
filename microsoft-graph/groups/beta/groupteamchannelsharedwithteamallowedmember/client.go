package groupteamchannelsharedwithteamallowedmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamChannelSharedWithTeamAllowedMemberClient struct {
	Client *msgraph.Client
}

func NewGroupTeamChannelSharedWithTeamAllowedMemberClientWithBaseURI(api sdkEnv.Api) (*GroupTeamChannelSharedWithTeamAllowedMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamchannelsharedwithteamallowedmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamChannelSharedWithTeamAllowedMemberClient: %+v", err)
	}

	return &GroupTeamChannelSharedWithTeamAllowedMemberClient{
		Client: client,
	}, nil
}
