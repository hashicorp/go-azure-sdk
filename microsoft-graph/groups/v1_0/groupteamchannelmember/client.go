package groupteamchannelmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamChannelMemberClient struct {
	Client *msgraph.Client
}

func NewGroupTeamChannelMemberClientWithBaseURI(api sdkEnv.Api) (*GroupTeamChannelMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamchannelmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamChannelMemberClient: %+v", err)
	}

	return &GroupTeamChannelMemberClient{
		Client: client,
	}, nil
}
