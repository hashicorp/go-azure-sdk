package mejoinedteamprimarychannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamPrimaryChannelClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamPrimaryChannelClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamPrimaryChannelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamprimarychannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamPrimaryChannelClient: %+v", err)
	}

	return &MeJoinedTeamPrimaryChannelClient{
		Client: client,
	}, nil
}
