package conversationthreadpostmention

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConversationThreadPostMentionClient struct {
	Client *msgraph.Client
}

func NewConversationThreadPostMentionClientWithBaseURI(sdkApi sdkEnv.Api) (*ConversationThreadPostMentionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conversationthreadpostmention", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConversationThreadPostMentionClient: %+v", err)
	}

	return &ConversationThreadPostMentionClient{
		Client: client,
	}, nil
}
