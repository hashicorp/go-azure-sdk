// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_11_01_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/iotcentral/2021-11-01-preview/apps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/iotcentral/2021-11-01-preview/networking"
)

type Client struct {
	Apps       *apps.AppsClient
	Networking *networking.NetworkingClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	appsClient := apps.NewAppsClientWithBaseURI(endpoint)
	configureAuthFunc(&appsClient.Client)

	networkingClient := networking.NewNetworkingClientWithBaseURI(endpoint)
	configureAuthFunc(&networkingClient.Client)

	return Client{
		Apps:       &appsClient,
		Networking: &networkingClient,
	}
}
