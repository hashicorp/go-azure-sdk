package conversationthreadpost

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConversationThreadPostClient struct {
	Client *msgraph.Client
}

func NewConversationThreadPostClientWithBaseURI(sdkApi sdkEnv.Api) (*ConversationThreadPostClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conversationthreadpost", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConversationThreadPostClient: %+v", err)
	}

	return &ConversationThreadPostClient{
		Client: client,
	}, nil
}
