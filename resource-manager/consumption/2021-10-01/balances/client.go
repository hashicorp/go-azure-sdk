package balances

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BalancesClient struct {
	Client *resourcemanager.Client
}

func NewBalancesClientWithBaseURI(sdkApi sdkEnv.Api) (*BalancesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "balances", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BalancesClient: %+v", err)
	}

	return &BalancesClient{
		Client: client,
	}, nil
}
