package userjoinedteamchannelsharedwithteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamChannelSharedWithTeamClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamChannelSharedWithTeamClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamChannelSharedWithTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamchannelsharedwithteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamChannelSharedWithTeamClient: %+v", err)
	}

	return &UserJoinedTeamChannelSharedWithTeamClient{
		Client: client,
	}, nil
}
