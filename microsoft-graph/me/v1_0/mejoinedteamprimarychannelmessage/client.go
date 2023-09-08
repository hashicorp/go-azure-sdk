package mejoinedteamprimarychannelmessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamPrimaryChannelMessageClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamPrimaryChannelMessageClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamPrimaryChannelMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamprimarychannelmessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamPrimaryChannelMessageClient: %+v", err)
	}

	return &MeJoinedTeamPrimaryChannelMessageClient{
		Client: client,
	}, nil
}
