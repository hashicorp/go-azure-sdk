package mejoinedteamgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamGroupClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamGroupClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamGroupClient: %+v", err)
	}

	return &MeJoinedTeamGroupClient{
		Client: client,
	}, nil
}
