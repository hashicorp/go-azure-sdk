package v2023_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/configurationstores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/deletedconfigurationstores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/keyvalues"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/operations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/replicas"
)

type Client struct {
	ConfigurationStores        *configurationstores.ConfigurationStoresClient
	DeletedConfigurationStores *deletedconfigurationstores.DeletedConfigurationStoresClient
	KeyValues                  *keyvalues.KeyValuesClient
	Operations                 *operations.OperationsClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	Replicas                   *replicas.ReplicasClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	configurationStoresClient := configurationstores.NewConfigurationStoresClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationStoresClient.Client)

	deletedConfigurationStoresClient := deletedconfigurationstores.NewDeletedConfigurationStoresClientWithBaseURI(endpoint)
	configureAuthFunc(&deletedConfigurationStoresClient.Client)

	keyValuesClient := keyvalues.NewKeyValuesClientWithBaseURI(endpoint)
	configureAuthFunc(&keyValuesClient.Client)

	operationsClient := operations.NewOperationsClientWithBaseURI(endpoint)
	configureAuthFunc(&operationsClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	replicasClient := replicas.NewReplicasClientWithBaseURI(endpoint)
	configureAuthFunc(&replicasClient.Client)

	return Client{
		ConfigurationStores:        &configurationStoresClient,
		DeletedConfigurationStores: &deletedConfigurationStoresClient,
		KeyValues:                  &keyValuesClient,
		Operations:                 &operationsClient,
		PrivateEndpointConnections: &privateEndpointConnectionsClient,
		PrivateLinkResources:       &privateLinkResourcesClient,
		Replicas:                   &replicasClient,
	}
}
