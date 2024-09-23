package conversation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConversationClient struct {
	Client *msgraph.Client
}

func NewConversationClientWithBaseURI(sdkApi sdkEnv.Api) (*ConversationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conversation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConversationClient: %+v", err)
	}

	return &ConversationClient{
		Client: client,
	}, nil
}
