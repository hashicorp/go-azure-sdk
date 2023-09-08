package mejoinedteamtag

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamTagClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamTagClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamTagClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamtag", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamTagClient: %+v", err)
	}

	return &MeJoinedTeamTagClient{
		Client: client,
	}, nil
}
