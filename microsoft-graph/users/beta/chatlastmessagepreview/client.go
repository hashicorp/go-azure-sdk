package chatlastmessagepreview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatLastMessagePreviewClient struct {
	Client *msgraph.Client
}

func NewChatLastMessagePreviewClientWithBaseURI(sdkApi sdkEnv.Api) (*ChatLastMessagePreviewClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chatlastmessagepreview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChatLastMessagePreviewClient: %+v", err)
	}

	return &ChatLastMessagePreviewClient{
		Client: client,
	}, nil
}
