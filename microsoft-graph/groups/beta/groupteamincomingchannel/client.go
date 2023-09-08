package groupteamincomingchannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamIncomingChannelClient struct {
	Client *msgraph.Client
}

func NewGroupTeamIncomingChannelClientWithBaseURI(api sdkEnv.Api) (*GroupTeamIncomingChannelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamincomingchannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamIncomingChannelClient: %+v", err)
	}

	return &GroupTeamIncomingChannelClient{
		Client: client,
	}, nil
}
