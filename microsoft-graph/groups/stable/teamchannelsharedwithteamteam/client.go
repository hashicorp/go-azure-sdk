package teamchannelsharedwithteamteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelSharedWithTeamTeamClient struct {
	Client *msgraph.Client
}

func NewTeamChannelSharedWithTeamTeamClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelSharedWithTeamTeamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelsharedwithteamteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelSharedWithTeamTeamClient: %+v", err)
	}

	return &TeamChannelSharedWithTeamTeamClient{
		Client: client,
	}, nil
}
