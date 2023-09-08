package mechatmessagehostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeChatMessageHostedContentClient struct {
	Client *msgraph.Client
}

func NewMeChatMessageHostedContentClientWithBaseURI(api sdkEnv.Api) (*MeChatMessageHostedContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mechatmessagehostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeChatMessageHostedContentClient: %+v", err)
	}

	return &MeChatMessageHostedContentClient{
		Client: client,
	}, nil
}
