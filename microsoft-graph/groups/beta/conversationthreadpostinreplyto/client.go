package conversationthreadpostinreplyto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConversationThreadPostInReplyToClient struct {
	Client *msgraph.Client
}

func NewConversationThreadPostInReplyToClientWithBaseURI(sdkApi sdkEnv.Api) (*ConversationThreadPostInReplyToClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conversationthreadpostinreplyto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConversationThreadPostInReplyToClient: %+v", err)
	}

	return &ConversationThreadPostInReplyToClient{
		Client: client,
	}, nil
}
