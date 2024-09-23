package teamallchannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamAllChannelClient struct {
	Client *msgraph.Client
}

func NewTeamAllChannelClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamAllChannelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamallchannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamAllChannelClient: %+v", err)
	}

	return &TeamAllChannelClient{
		Client: client,
	}, nil
}
