package vaultusages

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VaultUsagesClient struct {
	Client *resourcemanager.Client
}

func NewVaultUsagesClientWithBaseURI(sdkApi sdkEnv.Api) (*VaultUsagesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "vaultusages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VaultUsagesClient: %+v", err)
	}

	return &VaultUsagesClient{
		Client: client,
	}, nil
}
