package groupteamchanneltab

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamChannelTabClient struct {
	Client *msgraph.Client
}

func NewGroupTeamChannelTabClientWithBaseURI(api sdkEnv.Api) (*GroupTeamChannelTabClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamchanneltab", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamChannelTabClient: %+v", err)
	}

	return &GroupTeamChannelTabClient{
		Client: client,
	}, nil
}
