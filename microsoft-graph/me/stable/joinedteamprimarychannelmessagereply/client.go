package joinedteamprimarychannelmessagereply

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPrimaryChannelMessageReplyClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPrimaryChannelMessageReplyClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPrimaryChannelMessageReplyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamprimarychannelmessagereply", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPrimaryChannelMessageReplyClient: %+v", err)
	}

	return &JoinedTeamPrimaryChannelMessageReplyClient{
		Client: client,
	}, nil
}
