package groupconversationthreadpostinreplytoextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupConversationThreadPostInReplyToExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupConversationThreadPostInReplyToExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupConversationThreadPostInReplyToExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupconversationthreadpostinreplytoextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupConversationThreadPostInReplyToExtensionClient: %+v", err)
	}

	return &GroupConversationThreadPostInReplyToExtensionClient{
		Client: client,
	}, nil
}
