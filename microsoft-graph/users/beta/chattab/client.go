package chattab

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatTabClient struct {
	Client *msgraph.Client
}

func NewChatTabClientWithBaseURI(sdkApi sdkEnv.Api) (*ChatTabClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chattab", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChatTabClient: %+v", err)
	}

	return &ChatTabClient{
		Client: client,
	}, nil
}
