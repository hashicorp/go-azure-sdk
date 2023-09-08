package groupconversationthreadpostmention

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupConversationThreadPostMentionClient struct {
	Client *msgraph.Client
}

func NewGroupConversationThreadPostMentionClientWithBaseURI(api sdkEnv.Api) (*GroupConversationThreadPostMentionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupconversationthreadpostmention", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupConversationThreadPostMentionClient: %+v", err)
	}

	return &GroupConversationThreadPostMentionClient{
		Client: client,
	}, nil
}
