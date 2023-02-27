// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_10_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceupdate/2022-10-01/deviceupdates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceupdate/2022-10-01/privateendpointconnectionproxies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceupdate/2022-10-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceupdate/2022-10-01/privatelinkresources"
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
