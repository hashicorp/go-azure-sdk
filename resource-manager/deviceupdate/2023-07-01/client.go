package v2023_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceupdate/2023-07-01/deviceupdates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceupdate/2023-07-01/privateendpointconnectionproxies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceupdate/2023-07-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceupdate/2023-07-01/privatelinkresources"
)

type Client struct {
	Deviceupdates                    *deviceupdates.DeviceupdatesClient
	PrivateEndpointConnectionProxies *privateendpointconnectionproxies.PrivateEndpointConnectionProxiesClient
	PrivateEndpointConnections       *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources             *privatelinkresources.PrivateLinkResourcesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	deviceupdatesClient := deviceupdates.NewDeviceupdatesClientWithBaseURI(endpoint)
	configureAuthFunc(&deviceupdatesClient.Client)

	privateEndpointConnectionProxiesClient := privateendpointconnectionproxies.NewPrivateEndpointConnectionProxiesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionProxiesClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	return Client{
		Deviceupdates:                    &deviceupdatesClient,
		PrivateEndpointConnectionProxies: &privateEndpointConnectionProxiesClient,
		PrivateEndpointConnections:       &privateEndpointConnectionsClient,
		PrivateLinkResources:             &privateLinkResourcesClient,
	}
}
