package v2024_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/azurefleet/2024-11-01/fleets"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Fleets *fleets.FleetsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	fleetsClient, err := fleets.NewFleetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Fleets client: %+v", err)
	}
	configureFunc(fleetsClient.Client)

	return &Client{
		Fleets: fleetsClient,
	}, nil
}
