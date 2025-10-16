package namespaceassets

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceAssetsClient struct {
	Client *resourcemanager.Client
}

func NewNamespaceAssetsClientWithBaseURI(sdkApi sdkEnv.Api) (*NamespaceAssetsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "namespaceassets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NamespaceAssetsClient: %+v", err)
	}

	return &NamespaceAssetsClient{
		Client: client,
	}, nil
}
