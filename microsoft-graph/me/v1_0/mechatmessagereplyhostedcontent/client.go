package mechatmessagereplyhostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeChatMessageReplyHostedContentClient struct {
	Client *msgraph.Client
}

func NewMeChatMessageReplyHostedContentClientWithBaseURI(api sdkEnv.Api) (*MeChatMessageReplyHostedContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mechatmessagereplyhostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeChatMessageReplyHostedContentClient: %+v", err)
	}

	return &MeChatMessageReplyHostedContentClient{
		Client: client,
	}, nil
}
