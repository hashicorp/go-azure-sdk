package extension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionClient struct {
	Client *msgraph.Client
}

func NewExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*ExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "extension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExtensionClient: %+v", err)
	}

	return &ExtensionClient{
		Client: client,
	}, nil
}
