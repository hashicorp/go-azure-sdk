package v2017_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2017-04-01/hybridconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2017-04-01/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2017-04-01/wcfrelays"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	HybridConnections *hybridconnections.HybridConnectionsClient
	Namespaces        *namespaces.NamespacesClient
	WCFRelays         *wcfrelays.WCFRelaysClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	hybridConnectionsClient, err := hybridconnections.NewHybridConnectionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building HybridConnections client: %+v", err)
	}
	configureFunc(hybridConnectionsClient.Client)

	namespacesClient, err := namespaces.NewNamespacesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Namespaces client: %+v", err)
	}
	configureFunc(namespacesClient.Client)

	wCFRelaysClient, err := wcfrelays.NewWCFRelaysClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building WCFRelays client: %+v", err)
	}
	configureFunc(wCFRelaysClient.Client)

	return &Client{
		HybridConnections: hybridConnectionsClient,
		Namespaces:        namespacesClient,
		WCFRelays:         wCFRelaysClient,
	}, nil
}
