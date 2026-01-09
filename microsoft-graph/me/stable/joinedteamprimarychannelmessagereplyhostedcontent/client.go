package joinedteamprimarychannelmessagereplyhostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPrimaryChannelMessageReplyHostedContentClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPrimaryChannelMessageReplyHostedContentClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPrimaryChannelMessageReplyHostedContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamprimarychannelmessagereplyhostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPrimaryChannelMessageReplyHostedContentClient: %+v", err)
	}

	return &JoinedTeamPrimaryChannelMessageReplyHostedContentClient{
		Client: client,
	}, nil
}
