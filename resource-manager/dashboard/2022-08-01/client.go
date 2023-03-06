package v2022_08_01

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2022-08-01/grafanaresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2022-08-01/privateendpointconnection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2022-08-01/privatelinkresource"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	GrafanaResource           *grafanaresource.GrafanaResourceClient
	PrivateEndpointConnection *privateendpointconnection.PrivateEndpointConnectionClient
	PrivateLinkResource       *privatelinkresource.PrivateLinkResourceClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	grafanaResourceClient, err := grafanaresource.NewGrafanaResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GrafanaResource client: %+v", err)
	}
	configureFunc(grafanaResourceClient.Client)

	privateEndpointConnectionClient, err := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnection client: %+v", err)
	}
	configureFunc(privateEndpointConnectionClient.Client)

	privateLinkResourceClient, err := privatelinkresource.NewPrivateLinkResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResource client: %+v", err)
	}
	configureFunc(privateLinkResourceClient.Client)

	return &Client{
		GrafanaResource:           grafanaResourceClient,
		PrivateEndpointConnection: privateEndpointConnectionClient,
		PrivateLinkResource:       privateLinkResourceClient,
	}, nil
}
