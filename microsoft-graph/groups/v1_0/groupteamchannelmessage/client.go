package groupteamchannelmessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamChannelMessageClient struct {
	Client *msgraph.Client
}

func NewGroupTeamChannelMessageClientWithBaseURI(api sdkEnv.Api) (*GroupTeamChannelMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamchannelmessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamChannelMessageClient: %+v", err)
	}

	return &GroupTeamChannelMessageClient{
		Client: client,
	}, nil
}
