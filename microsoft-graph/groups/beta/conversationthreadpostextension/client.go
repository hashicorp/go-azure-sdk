package conversationthreadpostextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConversationThreadPostExtensionClient struct {
	Client *msgraph.Client
}

func NewConversationThreadPostExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*ConversationThreadPostExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conversationthreadpostextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConversationThreadPostExtensionClient: %+v", err)
	}

	return &ConversationThreadPostExtensionClient{
		Client: client,
	}, nil
}
