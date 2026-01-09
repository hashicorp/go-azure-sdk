package chatmessagehostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageHostedContentClient struct {
	Client *msgraph.Client
}

func NewChatMessageHostedContentClientWithBaseURI(sdkApi sdkEnv.Api) (*ChatMessageHostedContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chatmessagehostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChatMessageHostedContentClient: %+v", err)
	}

	return &ChatMessageHostedContentClient{
		Client: client,
	}, nil
}
