package chatmessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageClient struct {
	Client *msgraph.Client
}

func NewChatMessageClientWithBaseURI(sdkApi sdkEnv.Api) (*ChatMessageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chatmessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChatMessageClient: %+v", err)
	}

	return &ChatMessageClient{
		Client: client,
	}, nil
}
