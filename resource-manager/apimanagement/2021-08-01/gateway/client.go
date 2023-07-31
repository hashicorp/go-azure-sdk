package gateway

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayClient struct {
	Client *resourcemanager.Client
}

func NewGatewayClientWithBaseURI(api sdkEnv.Api) (*GatewayClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "gateway", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GatewayClient: %+v", err)
	}

	return &GatewayClient{
		Client: client,
	}, nil
}
