package mejoinedteamchannelsharedwithteamteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamChannelSharedWithTeamTeamClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamChannelSharedWithTeamTeamClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamChannelSharedWithTeamTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamchannelsharedwithteamteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamChannelSharedWithTeamTeamClient: %+v", err)
	}

	return &MeJoinedTeamChannelSharedWithTeamTeamClient{
		Client: client,
	}, nil
}
