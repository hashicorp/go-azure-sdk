package threadpostinreplytoextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreadPostInReplyToExtensionClient struct {
	Client *msgraph.Client
}

func NewThreadPostInReplyToExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*ThreadPostInReplyToExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "threadpostinreplytoextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ThreadPostInReplyToExtensionClient: %+v", err)
	}

	return &ThreadPostInReplyToExtensionClient{
		Client: client,
	}, nil
}
