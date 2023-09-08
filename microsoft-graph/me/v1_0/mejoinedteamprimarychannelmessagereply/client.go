package mejoinedteamprimarychannelmessagereply

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamPrimaryChannelMessageReplyClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamPrimaryChannelMessageReplyClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamPrimaryChannelMessageReplyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamprimarychannelmessagereply", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamPrimaryChannelMessageReplyClient: %+v", err)
	}

	return &MeJoinedTeamPrimaryChannelMessageReplyClient{
		Client: client,
	}, nil
}
