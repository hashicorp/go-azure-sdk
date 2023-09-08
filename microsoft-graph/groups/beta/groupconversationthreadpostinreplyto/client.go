package groupconversationthreadpostinreplyto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupConversationThreadPostInReplyToClient struct {
	Client *msgraph.Client
}

func NewGroupConversationThreadPostInReplyToClientWithBaseURI(api sdkEnv.Api) (*GroupConversationThreadPostInReplyToClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupconversationthreadpostinreplyto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupConversationThreadPostInReplyToClient: %+v", err)
	}

	return &GroupConversationThreadPostInReplyToClient{
		Client: client,
	}, nil
}
