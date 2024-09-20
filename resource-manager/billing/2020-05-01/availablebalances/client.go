package availablebalances

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailableBalancesClient struct {
	Client *resourcemanager.Client
}

func NewAvailableBalancesClientWithBaseURI(sdkApi sdkEnv.Api) (*AvailableBalancesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "availablebalances", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AvailableBalancesClient: %+v", err)
	}

	return &AvailableBalancesClient{
		Client: client,
	}, nil
}
