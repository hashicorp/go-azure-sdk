package conversationthread

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConversationThreadClient struct {
	Client *msgraph.Client
}

func NewConversationThreadClientWithBaseURI(sdkApi sdkEnv.Api) (*ConversationThreadClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conversationthread", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConversationThreadClient: %+v", err)
	}

	return &ConversationThreadClient{
		Client: client,
	}, nil
}
