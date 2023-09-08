package mejoinedteamprimarychanneltab

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamPrimaryChannelTabClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamPrimaryChannelTabClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamPrimaryChannelTabClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamprimarychanneltab", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamPrimaryChannelTabClient: %+v", err)
	}

	return &MeJoinedTeamPrimaryChannelTabClient{
		Client: client,
	}, nil
}
