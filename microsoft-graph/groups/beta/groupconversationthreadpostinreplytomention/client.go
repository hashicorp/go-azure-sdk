package groupconversationthreadpostinreplytomention

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupConversationThreadPostInReplyToMentionClient struct {
	Client *msgraph.Client
}

func NewGroupConversationThreadPostInReplyToMentionClientWithBaseURI(api sdkEnv.Api) (*GroupConversationThreadPostInReplyToMentionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupconversationthreadpostinreplytomention", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupConversationThreadPostInReplyToMentionClient: %+v", err)
	}

	return &GroupConversationThreadPostInReplyToMentionClient{
		Client: client,
	}, nil
}
