package v2024_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2024-10-01/grafanaplugin"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2024-10-01/grafanaresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2024-10-01/integrationfabric"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2024-10-01/managedprivateendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2024-10-01/privateendpointconnection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2024-10-01/privatelinkresource"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	GrafanaPlugin             *grafanaplugin.GrafanaPluginClient
	GrafanaResource           *grafanaresource.GrafanaResourceClient
	IntegrationFabric         *integrationfabric.IntegrationFabricClient
	ManagedPrivateEndpoints   *managedprivateendpoints.ManagedPrivateEndpointsClient
	PrivateEndpointConnection *privateendpointconnection.PrivateEndpointConnectionClient
	PrivateLinkResource       *privatelinkresource.PrivateLinkResourceClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	grafanaPluginClient, err := grafanaplugin.NewGrafanaPluginClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GrafanaPlugin client: %+v", err)
	}
	configureFunc(grafanaPluginClient.Client)

	grafanaResourceClient, err := grafanaresource.NewGrafanaResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GrafanaResource client: %+v", err)
	}
	configureFunc(grafanaResourceClient.Client)

	integrationFabricClient, err := integrationfabric.NewIntegrationFabricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationFabric client: %+v", err)
	}
	configureFunc(integrationFabricClient.Client)

	managedPrivateEndpointsClient, err := managedprivateendpoints.NewManagedPrivateEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedPrivateEndpoints client: %+v", err)
	}
	configureFunc(managedPrivateEndpointsClient.Client)

	privateEndpointConnectionClient, err := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnection client: %+v", err)
	}
	configureFunc(privateEndpointConnectionClient.Client)

	privateLinkResourceClient, err := privatelinkresource.NewPrivateLinkResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResource client: %+v", err)
	}
	configureFunc(privateLinkResourceClient.Client)

	return &Client{
		GrafanaPlugin:             grafanaPluginClient,
		GrafanaResource:           grafanaResourceClient,
		IntegrationFabric:         integrationFabricClient,
		ManagedPrivateEndpoints:   managedPrivateEndpointsClient,
		PrivateEndpointConnection: privateEndpointConnectionClient,
		PrivateLinkResource:       privateLinkResourceClient,
	}, nil
}
