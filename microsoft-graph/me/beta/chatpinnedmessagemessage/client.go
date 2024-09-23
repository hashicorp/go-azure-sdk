package chatpinnedmessagemessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatPinnedMessageMessageClient struct {
	Client *msgraph.Client
}

func NewChatPinnedMessageMessageClientWithBaseURI(sdkApi sdkEnv.Api) (*ChatPinnedMessageMessageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chatpinnedmessagemessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChatPinnedMessageMessageClient: %+v", err)
	}

	return &ChatPinnedMessageMessageClient{
		Client: client,
	}, nil
}
