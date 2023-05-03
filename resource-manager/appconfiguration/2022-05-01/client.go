package v2022_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/configurationstores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/deletedconfigurationstores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/keyvalues"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/operations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ConfigurationStores        *configurationstores.ConfigurationStoresClient
	DeletedConfigurationStores *deletedconfigurationstores.DeletedConfigurationStoresClient
	KeyValues                  *keyvalues.KeyValuesClient
	Operations                 *operations.OperationsClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	configurationStoresClient, err := configurationstores.NewConfigurationStoresClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationStores client: %+v", err)
	}
	configureFunc(configurationStoresClient.Client)

	deletedConfigurationStoresClient, err := deletedconfigurationstores.NewDeletedConfigurationStoresClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DeletedConfigurationStores client: %+v", err)
	}
	configureFunc(deletedConfigurationStoresClient.Client)

	keyValuesClient, err := keyvalues.NewKeyValuesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building KeyValues client: %+v", err)
	}
	configureFunc(keyValuesClient.Client)

	operationsClient, err := operations.NewOperationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Operations client: %+v", err)
	}
	configureFunc(operationsClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	return &Client{
		ConfigurationStores:        configurationStoresClient,
		DeletedConfigurationStores: deletedConfigurationStoresClient,
		KeyValues:                  keyValuesClient,
		Operations:                 operationsClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
	}, nil
}
