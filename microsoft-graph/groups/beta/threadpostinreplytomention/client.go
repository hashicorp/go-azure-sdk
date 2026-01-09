package threadpostinreplytomention

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreadPostInReplyToMentionClient struct {
	Client *msgraph.Client
}

func NewThreadPostInReplyToMentionClientWithBaseURI(sdkApi sdkEnv.Api) (*ThreadPostInReplyToMentionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "threadpostinreplytomention", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ThreadPostInReplyToMentionClient: %+v", err)
	}

	return &ThreadPostInReplyToMentionClient{
		Client: client,
	}, nil
}
