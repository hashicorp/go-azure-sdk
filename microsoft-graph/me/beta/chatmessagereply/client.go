package chatmessagereply

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageReplyClient struct {
	Client *msgraph.Client
}

func NewChatMessageReplyClientWithBaseURI(sdkApi sdkEnv.Api) (*ChatMessageReplyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chatmessagereply", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChatMessageReplyClient: %+v", err)
	}

	return &ChatMessageReplyClient{
		Client: client,
	}, nil
}
