package groupteamchannelmessagehostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamChannelMessageHostedContentClient struct {
	Client *msgraph.Client
}

func NewGroupTeamChannelMessageHostedContentClientWithBaseURI(api sdkEnv.Api) (*GroupTeamChannelMessageHostedContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamchannelmessagehostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamChannelMessageHostedContentClient: %+v", err)
	}

	return &GroupTeamChannelMessageHostedContentClient{
		Client: client,
	}, nil
}
