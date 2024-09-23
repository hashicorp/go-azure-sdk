package threadpostextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreadPostExtensionClient struct {
	Client *msgraph.Client
}

func NewThreadPostExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*ThreadPostExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "threadpostextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ThreadPostExtensionClient: %+v", err)
	}

	return &ThreadPostExtensionClient{
		Client: client,
	}, nil
}
