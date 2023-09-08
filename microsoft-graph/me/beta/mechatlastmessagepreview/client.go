package mechatlastmessagepreview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeChatLastMessagePreviewClient struct {
	Client *msgraph.Client
}

func NewMeChatLastMessagePreviewClientWithBaseURI(api sdkEnv.Api) (*MeChatLastMessagePreviewClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mechatlastmessagepreview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeChatLastMessagePreviewClient: %+v", err)
	}

	return &MeChatLastMessagePreviewClient{
		Client: client,
	}, nil
}
