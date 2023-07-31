package gatewayhostnameconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayHostnameConfigurationClient struct {
	Client *resourcemanager.Client
}

func NewGatewayHostnameConfigurationClientWithBaseURI(api sdkEnv.Api) (*GatewayHostnameConfigurationClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "gatewayhostnameconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GatewayHostnameConfigurationClient: %+v", err)
	}

	return &GatewayHostnameConfigurationClient{
		Client: client,
	}, nil
}
