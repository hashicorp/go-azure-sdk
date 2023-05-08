package v2021_11_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/iotcentral/2021-11-01-preview/apps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/iotcentral/2021-11-01-preview/networking"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Apps       *apps.AppsClient
	Networking *networking.NetworkingClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	appsClient, err := apps.NewAppsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Apps client: %+v", err)
	}
	configureFunc(appsClient.Client)

	networkingClient, err := networking.NewNetworkingClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Networking client: %+v", err)
	}
	configureFunc(networkingClient.Client)

	return &Client{
		Apps:       appsClient,
		Networking: networkingClient,
	}, nil
}
