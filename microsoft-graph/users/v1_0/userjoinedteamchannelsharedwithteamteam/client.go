package userjoinedteamchannelsharedwithteamteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamChannelSharedWithTeamTeamClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamChannelSharedWithTeamTeamClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamChannelSharedWithTeamTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamchannelsharedwithteamteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamChannelSharedWithTeamTeamClient: %+v", err)
	}

	return &UserJoinedTeamChannelSharedWithTeamTeamClient{
		Client: client,
	}, nil
}
