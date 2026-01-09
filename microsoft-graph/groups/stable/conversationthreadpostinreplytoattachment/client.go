package conversationthreadpostinreplytoattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConversationThreadPostInReplyToAttachmentClient struct {
	Client *msgraph.Client
}

func NewConversationThreadPostInReplyToAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*ConversationThreadPostInReplyToAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conversationthreadpostinreplytoattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConversationThreadPostInReplyToAttachmentClient: %+v", err)
	}

	return &ConversationThreadPostInReplyToAttachmentClient{
		Client: client,
	}, nil
}
