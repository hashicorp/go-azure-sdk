package vaultextendedinfo

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VaultExtendedInfoClient struct {
	Client *resourcemanager.Client
}

func NewVaultExtendedInfoClientWithBaseURI(sdkApi sdkEnv.Api) (*VaultExtendedInfoClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "vaultextendedinfo", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VaultExtendedInfoClient: %+v", err)
	}

	return &VaultExtendedInfoClient{
		Client: client,
	}, nil
}
