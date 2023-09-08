package groupconversationthreadpostattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupConversationThreadPostAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupConversationThreadPostAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupConversationThreadPostAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupconversationthreadpostattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupConversationThreadPostAttachmentClient: %+v", err)
	}

	return &GroupConversationThreadPostAttachmentClient{
		Client: client,
	}, nil
}
