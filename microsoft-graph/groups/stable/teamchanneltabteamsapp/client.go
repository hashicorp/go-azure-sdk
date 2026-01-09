package teamchanneltabteamsapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelTabTeamsAppClient struct {
	Client *msgraph.Client
}

func NewTeamChannelTabTeamsAppClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelTabTeamsAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchanneltabteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelTabTeamsAppClient: %+v", err)
	}

	return &TeamChannelTabTeamsAppClient{
		Client: client,
	}, nil
}
