package v2023_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/fabric/2023-11-01/fabriccapacities"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	FabricCapacities *fabriccapacities.FabricCapacitiesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	fabricCapacitiesClient, err := fabriccapacities.NewFabricCapacitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FabricCapacities client: %+v", err)
	}
	configureFunc(fabricCapacitiesClient.Client)

	return &Client{
		FabricCapacities: fabricCapacitiesClient,
	}, nil
}
