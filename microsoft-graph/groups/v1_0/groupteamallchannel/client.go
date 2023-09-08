package groupteamallchannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamAllChannelClient struct {
	Client *msgraph.Client
}

func NewGroupTeamAllChannelClientWithBaseURI(api sdkEnv.Api) (*GroupTeamAllChannelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamallchannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamAllChannelClient: %+v", err)
	}

	return &GroupTeamAllChannelClient{
		Client: client,
	}, nil
}
