// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2017_04_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2017-04-01/hybridconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2017-04-01/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2017-04-01/wcfrelays"
)

type Client struct {
	HybridConnections *hybridconnections.HybridConnectionsClient
	Namespaces        *namespaces.NamespacesClient
	WCFRelays         *wcfrelays.WCFRelaysClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	hybridConnectionsClient := hybridconnections.NewHybridConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&hybridConnectionsClient.Client)

	namespacesClient := namespaces.NewNamespacesClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesClient.Client)

	wCFRelaysClient := wcfrelays.NewWCFRelaysClientWithBaseURI(endpoint)
	configureAuthFunc(&wCFRelaysClient.Client)

	return Client{
		HybridConnections: &hybridConnectionsClient,
		Namespaces:        &namespacesClient,
		WCFRelays:         &wCFRelaysClient,
	}
}
