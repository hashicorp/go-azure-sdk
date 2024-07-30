package chatmessagereplyhostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageReplyHostedContentClient struct {
	Client *msgraph.Client
}

func NewChatMessageReplyHostedContentClientWithBaseURI(sdkApi sdkEnv.Api) (*ChatMessageReplyHostedContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chatmessagereplyhostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChatMessageReplyHostedContentClient: %+v", err)
	}

	return &ChatMessageReplyHostedContentClient{
		Client: client,
	}, nil
}
