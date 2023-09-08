package groupteamchannelsharedwithteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamChannelSharedWithTeamClient struct {
	Client *msgraph.Client
}

func NewGroupTeamChannelSharedWithTeamClientWithBaseURI(api sdkEnv.Api) (*GroupTeamChannelSharedWithTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamchannelsharedwithteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamChannelSharedWithTeamClient: %+v", err)
	}

	return &GroupTeamChannelSharedWithTeamClient{
		Client: client,
	}, nil
}
