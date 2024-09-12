package chat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatClient struct {
	Client *msgraph.Client
}

func NewChatClientWithBaseURI(sdkApi sdkEnv.Api) (*ChatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChatClient: %+v", err)
	}

	return &ChatClient{
		Client: client,
	}, nil
}
