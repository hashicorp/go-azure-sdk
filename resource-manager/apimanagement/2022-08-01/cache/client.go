package cache

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CacheClient struct {
	Client *resourcemanager.Client
}

func NewCacheClientWithBaseURI(api environments.Api) (*CacheClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "cache", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CacheClient: %+v", err)
	}

	return &CacheClient{
		Client: client,
	}, nil
}
