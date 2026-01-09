package messageextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageExtensionClient struct {
	Client *msgraph.Client
}

func NewMessageExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*MessageExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "messageextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MessageExtensionClient: %+v", err)
	}

	return &MessageExtensionClient{
		Client: client,
	}, nil
}
