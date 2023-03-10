package v2021_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2021-11-01/hybridconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2021-11-01/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2021-11-01/namespacesprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2021-11-01/namespacesprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2021-11-01/wcfrelays"
)

type Client struct {
	HybridConnections                    *hybridconnections.HybridConnectionsClient
	Namespaces                           *namespaces.NamespacesClient
	NamespacesPrivateEndpointConnections *namespacesprivateendpointconnections.NamespacesPrivateEndpointConnectionsClient
	NamespacesPrivateLinkResources       *namespacesprivatelinkresources.NamespacesPrivateLinkResourcesClient
	WCFRelays                            *wcfrelays.WCFRelaysClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	hybridConnectionsClient := hybridconnections.NewHybridConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&hybridConnectionsClient.Client)

	namespacesClient := namespaces.NewNamespacesClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesClient.Client)

	namespacesPrivateEndpointConnectionsClient := namespacesprivateendpointconnections.NewNamespacesPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesPrivateEndpointConnectionsClient.Client)

	namespacesPrivateLinkResourcesClient := namespacesprivatelinkresources.NewNamespacesPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesPrivateLinkResourcesClient.Client)

	wCFRelaysClient := wcfrelays.NewWCFRelaysClientWithBaseURI(endpoint)
	configureAuthFunc(&wCFRelaysClient.Client)

	return Client{
		HybridConnections:                    &hybridConnectionsClient,
		Namespaces:                           &namespacesClient,
		NamespacesPrivateEndpointConnections: &namespacesPrivateEndpointConnectionsClient,
		NamespacesPrivateLinkResources:       &namespacesPrivateLinkResourcesClient,
		WCFRelays:                            &wCFRelaysClient,
	}
}
