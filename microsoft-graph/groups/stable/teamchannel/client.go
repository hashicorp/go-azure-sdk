package teamchannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelClient struct {
	Client *msgraph.Client
}

func NewTeamChannelClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelClient: %+v", err)
	}

	return &TeamChannelClient{
		Client: client,
	}, nil
}
