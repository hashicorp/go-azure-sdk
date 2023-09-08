package mechatmessagereply

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeChatMessageReplyClient struct {
	Client *msgraph.Client
}

func NewMeChatMessageReplyClientWithBaseURI(api sdkEnv.Api) (*MeChatMessageReplyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mechatmessagereply", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeChatMessageReplyClient: %+v", err)
	}

	return &MeChatMessageReplyClient{
		Client: client,
	}, nil
}
