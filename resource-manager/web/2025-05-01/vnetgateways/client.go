package vnetgateways

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VnetGatewaysClient struct {
	Client *resourcemanager.Client
}

func NewVnetGatewaysClientWithBaseURI(sdkApi sdkEnv.Api) (*VnetGatewaysClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "vnetgateways", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VnetGatewaysClient: %+v", err)
	}

	return &VnetGatewaysClient{
		Client: client,
	}, nil
}
