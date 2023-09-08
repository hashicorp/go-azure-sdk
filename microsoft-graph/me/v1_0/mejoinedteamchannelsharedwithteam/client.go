package mejoinedteamchannelsharedwithteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamChannelSharedWithTeamClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamChannelSharedWithTeamClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamChannelSharedWithTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamchannelsharedwithteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamChannelSharedWithTeamClient: %+v", err)
	}

	return &MeJoinedTeamChannelSharedWithTeamClient{
		Client: client,
	}, nil
}
