package v2022_08_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2022-08-01/grafanaresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2022-08-01/privateendpointconnection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2022-08-01/privatelinkresource"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	GrafanaResource           *grafanaresource.GrafanaResourceClient
	PrivateEndpointConnection *privateendpointconnection.PrivateEndpointConnectionClient
	PrivateLinkResource       *privatelinkresource.PrivateLinkResourceClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	grafanaResourceClient := grafanaresource.NewGrafanaResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&grafanaResourceClient.Client)

	privateEndpointConnectionClient := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionClient.Client)

	privateLinkResourceClient := privatelinkresource.NewPrivateLinkResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourceClient.Client)

	return Client{
		GrafanaResource:           &grafanaResourceClient,
		PrivateEndpointConnection: &privateEndpointConnectionClient,
		PrivateLinkResource:       &privateLinkResourceClient,
	}
}
