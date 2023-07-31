package gatewaylistkeys

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayListKeysClient struct {
	Client *resourcemanager.Client
}

func NewGatewayListKeysClientWithBaseURI(api environments.Api) (*GatewayListKeysClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "gatewaylistkeys", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GatewayListKeysClient: %+v", err)
	}

	return &GatewayListKeysClient{
		Client: client,
	}, nil
}
