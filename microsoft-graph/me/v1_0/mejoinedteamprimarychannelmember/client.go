package mejoinedteamprimarychannelmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamPrimaryChannelMemberClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamPrimaryChannelMemberClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamPrimaryChannelMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamprimarychannelmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamPrimaryChannelMemberClient: %+v", err)
	}

	return &MeJoinedTeamPrimaryChannelMemberClient{
		Client: client,
	}, nil
}
