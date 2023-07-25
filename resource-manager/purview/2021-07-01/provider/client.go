package provider

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderClient struct {
	Client *resourcemanager.Client
}

func NewProviderClientWithBaseURI(api environments.Api) (*ProviderClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "provider", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProviderClient: %+v", err)
	}

	return &ProviderClient{
		Client: client,
	}, nil
}
