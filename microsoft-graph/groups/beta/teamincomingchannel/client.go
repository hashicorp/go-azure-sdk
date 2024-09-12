package teamincomingchannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamIncomingChannelClient struct {
	Client *msgraph.Client
}

func NewTeamIncomingChannelClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamIncomingChannelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamincomingchannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamIncomingChannelClient: %+v", err)
	}

	return &TeamIncomingChannelClient{
		Client: client,
	}, nil
}
