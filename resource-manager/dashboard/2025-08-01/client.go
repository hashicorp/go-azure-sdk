package v2025_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2025-08-01/integrationfabrics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2025-08-01/manageddashboards"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2025-08-01/managedgrafanas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2025-08-01/managedprivateendpointmodels"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2025-08-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2025-08-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	IntegrationFabrics           *integrationfabrics.IntegrationFabricsClient
	ManagedDashboards            *manageddashboards.ManagedDashboardsClient
	ManagedGrafanas              *managedgrafanas.ManagedGrafanasClient
	ManagedPrivateEndpointModels *managedprivateendpointmodels.ManagedPrivateEndpointModelsClient
	PrivateEndpointConnections   *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources         *privatelinkresources.PrivateLinkResourcesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	integrationFabricsClient, err := integrationfabrics.NewIntegrationFabricsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationFabrics client: %+v", err)
	}
	configureFunc(integrationFabricsClient.Client)

	managedDashboardsClient, err := manageddashboards.NewManagedDashboardsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDashboards client: %+v", err)
	}
	configureFunc(managedDashboardsClient.Client)

	managedGrafanasClient, err := managedgrafanas.NewManagedGrafanasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedGrafanas client: %+v", err)
	}
	configureFunc(managedGrafanasClient.Client)

	managedPrivateEndpointModelsClient, err := managedprivateendpointmodels.NewManagedPrivateEndpointModelsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedPrivateEndpointModels client: %+v", err)
	}
	configureFunc(managedPrivateEndpointModelsClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	return &Client{
		IntegrationFabrics:           integrationFabricsClient,
		ManagedDashboards:            managedDashboardsClient,
		ManagedGrafanas:              managedGrafanasClient,
		ManagedPrivateEndpointModels: managedPrivateEndpointModelsClient,
		PrivateEndpointConnections:   privateEndpointConnectionsClient,
		PrivateLinkResources:         privateLinkResourcesClient,
	}, nil
}
