package hypervrunasaccounts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVRunAsAccountsClient struct {
	Client *resourcemanager.Client
}

func NewHyperVRunAsAccountsClientWithBaseURI(api sdkEnv.Api) (*HyperVRunAsAccountsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "hypervrunasaccounts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HyperVRunAsAccountsClient: %+v", err)
	}

	return &HyperVRunAsAccountsClient{
		Client: client,
	}, nil
}
