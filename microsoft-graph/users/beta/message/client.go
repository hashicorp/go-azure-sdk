package message

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageClient struct {
	Client *msgraph.Client
}

func NewMessageClientWithBaseURI(sdkApi sdkEnv.Api) (*MessageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "message", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MessageClient: %+v", err)
	}

	return &MessageClient{
		Client: client,
	}, nil
}
