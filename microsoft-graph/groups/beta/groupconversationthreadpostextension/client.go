package groupconversationthreadpostextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupConversationThreadPostExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupConversationThreadPostExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupConversationThreadPostExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupconversationthreadpostextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupConversationThreadPostExtensionClient: %+v", err)
	}

	return &GroupConversationThreadPostExtensionClient{
		Client: client,
	}, nil
}
