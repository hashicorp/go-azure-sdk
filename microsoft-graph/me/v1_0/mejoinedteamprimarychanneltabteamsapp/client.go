package mejoinedteamprimarychanneltabteamsapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamPrimaryChannelTabTeamsAppClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamPrimaryChannelTabTeamsAppClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamPrimaryChannelTabTeamsAppClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamprimarychanneltabteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamPrimaryChannelTabTeamsAppClient: %+v", err)
	}

	return &MeJoinedTeamPrimaryChannelTabTeamsAppClient{
		Client: client,
	}, nil
}
